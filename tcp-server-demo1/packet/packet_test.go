package packet

import "testing"

func TestDecode(t *testing.T) {
	s := &Submit{}
	err := s.Decode([]byte("12345678hello"))
	if err != nil {
		t.Errorf("want nil, actual %s", err.Error())
	}

	if s.ID != "12345678" {
		t.Errorf("want 12345678, actual %s", s.ID)
	}

	if string(s.Payload) != "hello" {
		t.Errorf("want hello, actual %s", string(s.Payload))
	}
}

func TestEncode(t *testing.T) {
	s := &Submit{
		ID:      "12345678",
		Payload: []byte("hello"),
	}

	//fmt.Println(s)
	b, err := s.Encode()
	if err != nil {
		t.Errorf("want nil, actual %s", err.Error())
	}

	if string(b) != "12345678hello" {
		t.Errorf("want 12345678hello, actual %s", string(b))
	}

}
