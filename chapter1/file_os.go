package main

// import "io"

// func Walk(dir string, f func(b []byte, err error) error) error {
// 	files, err := os.ReadDir(dir)
// 	if err != nil { return err }
// 	for _, file := range files {
// 		path := filepath.Join(dir, file.Name())
// 		if !file.IsDir() {
// 			b, err := os.ReadFile(path)
// 			if err := f(b, err); err != nil { return err }
// 			continue
// 		}
// 		if err := Walk(path, f); err != nil { return err }
// 	}
// 	return nil
// }

// type REPL struct { stmts []string }
// func (r *REPL) Exec(w io.Writer, expr string) error {
// 	file, err := os.CreateTemp("", "repl*.go")
// 	if err != nil { return err }
// 	const src = `package main;import "fmt"; func main(){%s;fmt.Println(%s)}`
// 	fmt.Fprintf(file, src, strings.Join(r.stmts, ";"), expr)
// 	file.Close()
// 	defer os.RemoveAll(file.Name())
// 	cmd := exec.Command("go", "run", file.Name())
// 	cmd.Stdout, cmd.Stderr = w, io.Discard
// 	if err := cmd.Run(); err != nil { return err }
// 	return nil
// }

// func Vers(module string) ([]string, error) {
// 	dir, err := os.MkdirTemp("", "vers*")
// 	if err != nil { return nil, err }
// 	defer os.RemoveAll(dir)

// 	env := &exec.Env{Dir: dir}
// 	env.Run("go", "mod", "init", "tmpmodule")
// 	env.Run("go", "get", module)
// 	pr, pw := io.Pipe()
// 	go func() {
// 		env.Stdout = pw
// 		env.Run("go", "list", "-m", "-versions", "-json", module)
// 		pw.Close()
// 	}()

// 	var vers struct{ Versions []string }
// 	err = json.NewDecoder(pr).Decode(&vers)
// 	if err != nil { return nil, err }
// 	err = env.Err()
// 	if err != nil { return nil, err }
// 	return vers.Versions, nil
// }
