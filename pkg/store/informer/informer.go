package informer

import (
	"context"
	"sync"

	"go.etcd.io/etcd/clientv3"
	"go.etcd.io/etcd/mvcc/mvccpb"
)

// Informer
type Informer struct {
	channel   map[string]clientv3.WatchChan
	receivers map[string]*receiverList
}

type receiverList struct {
	list       []chan *Result
	usageCount uint64
	sync.RWMutex
}

// Start
func (inf *Informer) Start(stopCh chan struct{}) error {
	var err error
	for {
		select {
		case <-stopCh:
			return err
		}
	}
}

// DataStoreChExists
func (inf *Informer) DataStoreChExists(key string) bool {
	// TODO Iterate over "per directory /" map as we can have compacted the datastore channel
	// _, ok := inf.channel[key]
	return false
}

// Watch
func (inf *Informer) Watch(ctx context.Context, key string, dataStoreCh clientv3.WatchChan) (chan *Result, error) {
	if _, ok := inf.channel[key]; !ok {
		inf.channel[key] = dataStoreCh
		go inf.watch(ctx, key)
	}
	ch := make(chan *Result)
	inf.receivers[key].Lock()
	inf.receivers[key].list = append(inf.receivers[key].list, ch)
	inf.receivers[key].Unlock()
	return ch, nil
}

func (inf *Informer) watch(ctx context.Context, key string) {
	for {
		select {
		case <-ctx.Done():
			return
		case resp := <-inf.channel[key]:
			var errs []error
			if err := resp.Err(); err != nil {
				errs = append(errs, err)
			} else {
				errs = nil
			}
			for _, event := range resp.Events {
				result := &Result{
					Errors:  errs,
					Closed:  resp.Canceled,
					Name:    string(event.Kv.Key),
					State:   convertETCDtoResultState(event),
					Value:   string(event.Kv.Value),
					Version: event.Kv.Version,
				}

				inf.receivers[key].RLock()
				for _, receiver := range inf.receivers[key].list {
					// TODO should we lock the list here?
					receiver <- result
				}
				inf.receivers[key].RUnlock()
			}
			if resp.Canceled {
				// TODO we should probably log this
				return
			}
		}
	}
}

func convertETCDtoResultState(event *clientv3.Event) ResultState {
	if event.Type == mvccpb.DELETE {
		return ResultStateDeleted
	}
	if event.Type == mvccpb.PUT {
		if event.Kv.CreateRevision == event.Kv.Version {
			return ResultStateUpdated
		}
		return ResultStateCreated
	}
	return ResultStateUnknown
}
