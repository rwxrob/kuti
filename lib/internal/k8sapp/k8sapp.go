package k8sapp

import "log"

func RemoveSecrets(files ...string) error {
	log.Print(`would remove secrets from following files`)
	for _, name := range files {
		log.Println(name)
	}
	return nil
}
