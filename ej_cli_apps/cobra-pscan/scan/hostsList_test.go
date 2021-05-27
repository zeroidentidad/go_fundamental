package scan_test

import (
	"io/ioutil"
	"os"
	"testing"

	"pscan/scan"
)

func TestAdd(t *testing.T) {
	testCases := []struct {
		name         string
		host         string
		expectLen    int
		expectErrMsg string
	}{
		{"AddNew", "host2", 2, ""},
		{"AddExisting", "host1", 1, "Host host1 already in the list"},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			hl := &scan.HostsList{}

			// Inicializar lista
			if err := hl.Add("host1"); err != nil {
				t.Fatal(err)
			}

			err := hl.Add(tc.host)

			if tc.expectErrMsg != "" {
				if err == nil {
					t.Fatalf("Expected error, got nil instead\n")
				}

				if err.Error() != tc.expectErrMsg {
					t.Errorf("Expected error %q, got %q instead\n", tc.expectErrMsg, err)
				}

				return
			}

			if err != nil {
				t.Fatalf("Expected no error, got %q instead\n", err)
			}

			if len(hl.Hosts) != tc.expectLen {
				t.Errorf("Expected list length %d, got %d instead\n", tc.expectLen, len(hl.Hosts))
			}

			if hl.Hosts[1] != tc.host {
				t.Errorf("Expected host name %q as index 1, got %q instead\n", tc.host, hl.Hosts[1])
			}
		})
	}
}

func TestRemove(t *testing.T) {
	testCases := []struct {
		name         string
		host         string
		expectLen    int
		expectErrMsg string
	}{
		{"RemoveExisting", "host1", 1, ""},
		{"RemoveNotFound", "host3", 1, "Host host3 is not in the list"},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			hl := &scan.HostsList{}

			// Inicializar lista
			for _, h := range []string{"host1", "host2"} {
				if err := hl.Add(h); err != nil {
					t.Fatal(err)
				}
			}

			err := hl.Remove(tc.host)

			if tc.expectErrMsg != "" {
				if err == nil {
					t.Fatalf("Expected error, got nil instead\n")
				}

				if err.Error() != tc.expectErrMsg {
					t.Errorf("Expected error %q, got %q instead\n", tc.expectErrMsg, err)
				}

				return
			}

			if err != nil {
				t.Fatalf("Expected no error, got %q instead\n", err)
			}

			if len(hl.Hosts) != tc.expectLen {
				t.Errorf("Expected list length %d, got %d instead\n", tc.expectLen, len(hl.Hosts))
			}

			if hl.Hosts[0] == tc.host {
				t.Errorf("Host name %q should not be in the list\n", tc.host)
			}
		})
	}
}

func TestSaveLoad(t *testing.T) {
	hl1 := scan.HostsList{}
	hl2 := scan.HostsList{}

	hostName := "host1"
	hl1.Add(hostName)

	tf, err := ioutil.TempFile("", "")

	if err != nil {
		t.Fatalf("Error creating temp file: %s", err)

	}
	defer os.Remove(tf.Name())

	if err := hl1.Save(tf.Name()); err != nil {
		t.Fatalf("Error saving list to file: %s", err)

	}

	if err := hl2.Load(tf.Name()); err != nil {
		t.Fatalf("Error getting list from file: %s", err)

	}

	if hl1.Hosts[0] != hl2.Hosts[0] {
		t.Errorf("Host %q should match %q host.", hl1.Hosts[0], hl2.Hosts[0])
	}
}

func TestLoadNoFile(t *testing.T) {
	tf, err := ioutil.TempFile("", "")

	if err != nil {
		t.Fatalf("Error creating temp file: %s", err)
	}

	if err := os.Remove(tf.Name()); err != nil {
		t.Fatalf("Error deleting temp file: %s", err)
	}

	hl := &scan.HostsList{}

	if err := hl.Load(tf.Name()); err != nil {
		t.Errorf("Expected no error, got %q instead\n", err)
	}
}
