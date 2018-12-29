BEGIN { err = 1; n = split(strings, lines, /\n/); i = 1; }
lines[i] == $0 { if (i++ == n) err = 0; next; }
{ i = 1; }
END { exit err; }
