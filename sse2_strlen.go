package main

/*
#include <stdint.h>
#include <stddef.h>
#include <stdio.h>
#include <emmintrin.h>

// Source: https://porg.es/blog/ridiculous-utf-8-character-counting

#define GetMask(x) __builtin_ia32_pmovmskb128(x)
#define LoadBytes(x) __builtin_ia32_loaddqu(x)
#define CompareEquality(x,y) __builtin_ia32_pcmpeqb128((x),(y))
#define NotExpected(x) __builtin_expect((x),0)
#define And(x,y) __builtin_ia32_pand128((x),(y))

const char mask[16] = {
    0xc0, 0xc0, 0xc0, 0xc0,
    0xc0, 0xc0, 0xc0, 0xc0,
    0xc0, 0xc0, 0xc0, 0xc0,
    0xc0, 0xc0, 0xc0, 0xc0
};
const char match[16] = {
    0x80, 0x80, 0x80, 0x80,
    0x80, 0x80, 0x80, 0x80,
    0x80, 0x80, 0x80, 0x80,
    0x80, 0x80, 0x80, 0x80
};
const char zero[16] = { 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0 };
unsigned char HammingWeight[65536]; //initialized elsewhere

static size_t cp_strlen_utf8_sse2(const char *_s)
{
    const char *s;
    const __v16qi allZero = LoadBytes(zero);
    const __v16qi masking = LoadBytes(mask);
    const __v16qi matching = LoadBytes(match);

    __v16qi row;
    size_t count = 0;
    unsigned char b;

    // unaligned bytes
    for (s = _s; (uintptr_t) (s) & (sizeof(__v16qi) - 1); s++) {
        b = *s;
        if (b == '\0')
            goto done;
        count += (b >> 7) & ((~b) >> 6);
    }

    // Handle complete blocks.
    for (;; s += sizeof(__v16qi)) {
        // Prefetch
        __builtin_prefetch(&s[256], 0, 0);

        // Load Bytes
        row = LoadBytes(s);

        // Expect this to be false :)
        if (NotExpected(GetMask(
                                   // Check for zero bytes
                                      CompareEquality(allZero, row))))
            break;

        // Count number of non-starter bytes

        row = (__v16qi) And((__v2di) row, (__v2di) masking);
        row = CompareEquality(row, matching);
        count += HammingWeight[GetMask(row)];
    }

    //leftover bytes
    for (;; s++) {
        b = *s;
        if (b == '\0')
            break;
        count += (b >> 7) & ((~b) >> 6);
    }

  done:
    return ((s - _s) - count);
}*/
import "C"
import "unsafe"
import "github.com/kitech/php-go/phpgo"

func module_strlen(str string) int {
    c_str := C.CString(str)
    defer C.free(unsafe.Pointer(c_str))

    return int(C.cp_strlen_utf8_sse2(c_str))
}


func init() {
    phpgo.InitExtension("sse2_strlen", "0.0.1")
    phpgo.AddFunc("sse2_strlen", module_strlen)
}

func main() {
}
