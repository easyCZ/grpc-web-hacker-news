package firego_test

import (
	"log"
	"time"

	"github.com/zabawaba99/firego"
)

func ExampleFirebase_Auth() {
	fb := firego.New("https://someapp.firebaseio.com", nil)
	fb.Auth("my-token")
}

func ExampleFirebase_Child() {
	fb := firego.New("https://someapp.firebaseio.com", nil)
	childFB := fb.Child("some/child/path")

	log.Printf("My new ref %s\n", childFB)
}

func ExampleFirebase_Shallow() {
	fb := firego.New("https://someapp.firebaseio.com", nil)
	// Set value
	fb.Shallow(true)
	// Remove query parameter
	fb.Shallow(false)
}

func ExampleFirebase_IncludePriority() {
	fb := firego.New("https://someapp.firebaseio.com", nil)
	// Set value
	fb.IncludePriority(true)
	// Remove query parameter
	fb.IncludePriority(false)
}

func ExampleFirebase_StartAt() {
	fb := firego.New("https://someapp.firebaseio.com", nil)
	// Set value
	fb = fb.StartAt("a")
	// Remove query parameter
	fb = fb.StartAt("")
}

func ExampleFirebase_EndAt() {
	fb := firego.New("https://someapp.firebaseio.com", nil)
	// Set value
	fb = fb.EndAt("a")
	// Remove query parameter
	fb = fb.EndAt("")
}

func ExampleFirebase_OrderBy() {
	fb := firego.New("https://someapp.firebaseio.com", nil)
	// Set value
	fb = fb.OrderBy("a")
	// Remove query parameter
	fb = fb.OrderBy("")
}

func ExampleFirebase_LimitToFirst() {
	fb := firego.New("https://someapp.firebaseio.com", nil)
	// Set value
	fb = fb.LimitToFirst(5)
	// Remove query parameter
	fb = fb.LimitToFirst(-1)
}

func ExampleFirebase_LimitToLast() {
	fb := firego.New("https://someapp.firebaseio.com", nil)
	// Set value
	fb = fb.LimitToLast(8)
	// Remove query parameter
	fb = fb.LimitToLast(-1)
}

func ExampleFirebase_Push() {
	fb := firego.New("https://someapp.firebaseio.com", nil)
	newRef, err := fb.Push("my-value")
	if err != nil {
		log.Fatal(err)
	}

	log.Printf("My new ref %s\n", newRef)
}

func ExampleFirebase_Remove() {
	fb := firego.New("https://someapp.firebaseio.com/some/value", nil)
	if err := fb.Remove(); err != nil {
		log.Fatal(err)
	}
}

func ExampleFirebase_Set() {
	fb := firego.New("https://someapp.firebaseio.com", nil)

	v := map[string]interface{}{
		"foo": "bar",
		"bar": 1,
		"bez": []string{"hello", "world"},
	}
	if err := fb.Set(v); err != nil {
		log.Fatal(err)
	}
}

func ExampleFirebase_Update() {
	fb := firego.New("https://someapp.firebaseio.com/some/value", nil)
	if err := fb.Update("new-value"); err != nil {
		log.Fatal(err)
	}
}

func ExampleFirebase_Value() {
	fb := firego.New("https://someapp.firebaseio.com/some/value", nil)
	var v interface{}
	if err := fb.Value(v); err != nil {
		log.Fatal(err)
	}

	log.Printf("My value %v\n", v)
}

func ExampleFirebase_Watch() {
	fb := firego.New("https://someapp.firebaseio.com/some/value", nil)
	notifications := make(chan firego.Event)
	if err := fb.Watch(notifications); err != nil {
		log.Fatal(err)
	}

	for event := range notifications {
		log.Println("Event Received")
		log.Printf("Type: %s\n", event.Type)
		log.Printf("Path: %s\n", event.Path)
		log.Printf("Data: %v\n", event.Data)
		if event.Type == firego.EventTypeError {
			log.Print("Error occurred, loop ending")
		}
	}
}

func ExampleFirebase_StopWatching() {
	fb := firego.New("https://someapp.firebaseio.com/some/value", nil)
	notifications := make(chan firego.Event)
	if err := fb.Watch(notifications); err != nil {
		log.Fatal(err)
	}

	go func() {
		for range notifications {
		}
		log.Println("Channel closed")
	}()
	time.Sleep(10 * time.Millisecond) // let go routine start

	fb.StopWatching()
}
