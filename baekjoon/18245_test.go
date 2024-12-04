package baekjoon_test

import(
    "testing"
    "algorithm/baekjoon"
)

func TestMain(t *testing.T) {
    tests := []struct {
        input string
        step int
        expected string
    } {
        {"HZAOPAPCYSUENCBOINRDTCHODNATY", 2, "HAPPYUNBIRTHDAY"},
        {"RAEBDCVDEELFVGEHT", 2, "REDVELVET"},
        {"SIJEKLUMNLOPGQRI", 3, "SEULGI"},
        {"ISTURVWXEYZANBCDE", 4, "IRENE"},
        {"WEFGHEIJKLNMNOPDQRSTY", 5, "WENDY"},
        {"YUVWXYEZABCDREFGHII", 6, "YERI"},
        {"JJKLMNOOPQRSTUY", 7, "JOY"},
    }

    for _, tt := range tests {
        result := baekjoon.Decode(tt.input, tt.step)
        if result != tt.expected {
            t.Errorf("decode(%s, %d) = %s; want %s", tt.input, tt.step, result, tt.expected)
        }
    }
}
