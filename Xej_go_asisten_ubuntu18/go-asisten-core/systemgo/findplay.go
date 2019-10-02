package systemgo

import (
	"errors"
	"fmt"
	"io/ioutil"
	"strings"

	"../osfunc"
)

// FindPlay func
func (ctx *SystemGoModel) FindPlay(datos ...string) error {
	if len(datos) != 3 {
		return errors.New("Arguments Not Complete")
	}
	findplay(datos[0], datos[1], datos[2], ctx)
	return nil
}

// FindPlayEvent func
func (ctx *SystemGoModel) FindPlayEvent(datos ...string) {
	ctx.FindPlayEvents <- datos[0]
}

func findplay(play string, dir string, find string, ctx *SystemGoModel) error {
	var forPlay func(string) bool
	forPlay = func(dirf string) bool {
		files, err := ioutil.ReadDir(dirf)
		if err != nil {
			return false
		}
		for _, vals := range files {
			if vals.IsDir() == false {
				if strings.Contains(vals.Name(), find) {
					//log.Println("Encontrado: ", play+" "+dir+vals.Name())
					done := make(chan error, 1)
					cmd := osfunc.Setcommand(play + " \"" + dirf + vals.Name() + "\"")
					cmd.Start()
					go func() {
						done <- cmd.Wait()
					}()
					select {
					case event := <-ctx.FindPlayEvents:
						if event == "kill" {
							cmd.Process.Kill()
							return false
						} else if event == "continue" {
							cmd.Process.Kill()
							continue
						}
					case err := <-done:
						fmt.Println("Continue", err)
						continue
					}
					//osfunc.CommandRun(play + " \"" + dirf + vals.Name() + "\"")
				}
			} else {
				if forPlay(dirf+vals.Name()+"/") == false {
					return false
				}
			}
		}
		return true
	}
	forPlay(dir)
	return nil
}
