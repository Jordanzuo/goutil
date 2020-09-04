package stringUtil

import (
	"testing"
)

func TestSimilarity(t *testing.T) {
	source := ""
	target := ""
	expectedDistance := 0
	expectedSimilarity := 0.0
	gotDistance, gotSimilarity := Similarity(source, target)
	if gotDistance != expectedDistance {
		t.Errorf("Expected to get %d, now got %d", expectedDistance, gotDistance)
		return
	}
	if gotSimilarity != expectedSimilarity {
		t.Errorf("Expected to get %f, now got %f", expectedSimilarity, gotSimilarity)
		return
	}

	source = "Hello"
	target = ""
	expectedDistance = 5
	expectedSimilarity = 0.0
	gotDistance, gotSimilarity = Similarity(source, target)
	if gotDistance != expectedDistance {
		t.Errorf("Expected to get %d, now got %d", expectedDistance, gotDistance)
		return
	}
	if gotSimilarity != expectedSimilarity {
		t.Errorf("Expected to get %f, now got %f", expectedSimilarity, gotSimilarity)
		return
	}

	source = ""
	target = "Hello"
	expectedDistance = 5
	expectedSimilarity = 0.0
	gotDistance, gotSimilarity = Similarity(source, target)
	if gotDistance != expectedDistance {
		t.Errorf("Expected to get %d, now got %d", expectedDistance, gotDistance)
		return
	}
	if gotSimilarity != expectedSimilarity {
		t.Errorf("Expected to get %f, now got %f", expectedSimilarity, gotSimilarity)
		return
	}

	source = "Helo"
	target = "Hello"
	expectedDistance = 1
	expectedSimilarity = 4.0 / 5.0
	gotDistance, gotSimilarity = Similarity(source, target)
	if gotDistance != expectedDistance {
		t.Errorf("Expected to get %d, now got %d", expectedDistance, gotDistance)
		return
	}
	if gotSimilarity != expectedSimilarity {
		t.Errorf("Expected to get %f, now got %f", expectedSimilarity, gotSimilarity)
		return
	}

	source = "kitten"
	target = "sitten"
	expectedDistance = 1
	expectedSimilarity = 5.0 / 6.0
	gotDistance, gotSimilarity = Similarity(source, target)
	if gotDistance != expectedDistance {
		t.Errorf("Expected to get %d, now got %d", expectedDistance, gotDistance)
		return
	}
	if gotSimilarity != expectedSimilarity {
		t.Errorf("Expected to get %f, now got %f", expectedSimilarity, gotSimilarity)
		return
	}

	source = "Michael Jordan"
	target = "Michael Jordan"
	expectedDistance = 0
	expectedSimilarity = 1
	gotDistance, gotSimilarity = Similarity(source, target)
	if gotDistance != expectedDistance {
		t.Errorf("Expected to get %d, now got %d", expectedDistance, gotDistance)
		return
	}
	if gotSimilarity != expectedSimilarity {
		t.Errorf("Expected to get %f, now got %f", expectedSimilarity, gotSimilarity)
		return
	}
}
