package wait

import (
	"fmt"
	"log"
	"net/http"
	"time"
)

// WaitForServer attempts to contact the server of a url;
// it tries for 1 minute, using exponential backoff and reports an error if all attempts fail;
func WaitForServer(url string) error {

	const timeout = 1 * time.Minute
	deadline := time.Now().Add(timeout)

	for tries := 0; time.Now().Before(deadline); tries++ {
		_, err := http.Head(url)
		if err == nil {
			return nil // success
		}
		log.Printf("server not responding (%s); retrying...", err)
		time.Sleep(time.Second << uint(tries)) // exponential backoff
	}

	return fmt.Errorf("server %s failed to respond after %s", url, timeout)

}
