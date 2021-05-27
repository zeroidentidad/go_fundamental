// Package scan proporciona tipos y funciones para realizar
// escaneos de puertos TCP en una lista de hosts
package scan

import (
	"bufio"
	"fmt"
	"io/ioutil"
	"os"
	"sort"
)

// La lista de hosts representa una lista de hosts
// para ejecutar el escaneo de puertos
type HostsList struct {
	Hosts []string
}

// search busca hosts en la lista
func (hl *HostsList) search(host string) (bool, int) {
	sort.Slice(hl.Hosts, func(i, j int) bool {
		return hl.Hosts[i] < hl.Hosts[j]
	})

	i := sort.SearchStrings(hl.Hosts, host)
	if i < len(hl.Hosts) && hl.Hosts[i] == host {
		return true, i
	}

	return false, -1
}

// Add agrega un host a la lista
func (hl *HostsList) Add(host string) error {
	if found, _ := hl.search(host); found {
		return fmt.Errorf("Host %s already in the list", host)
	}

	hl.Hosts = append(hl.Hosts, host)
	return nil
}

// Remove elimina un host de la lista
func (hl *HostsList) Remove(host string) error {
	if found, i := hl.search(host); found {
		hl.Hosts = append(hl.Hosts[:i], hl.Hosts[i+1:]...)
		return nil
	}

	return fmt.Errorf("Host %s is not in the list", host)
}

// Load obtiene hosts de un archivo de hosts
func (hl *HostsList) Load(hostsFile string) error {
	f, err := os.Open(hostsFile)
	if err != nil {
		if os.IsNotExist(err) {
			return nil
		}
		return err
	}
	defer f.Close()

	scanner := bufio.NewScanner(f)

	for scanner.Scan() {
		hl.Hosts = append(hl.Hosts, scanner.Text())
	}

	return nil
}

// Save guarda los hosts en un archivo de hosts
func (hl *HostsList) Save(hostsFile string) error {
	output := ""

	for _, h := range hl.Hosts {
		output += fmt.Sprintln(h)
	}

	return ioutil.WriteFile(hostsFile, []byte(output), 0644)
}
