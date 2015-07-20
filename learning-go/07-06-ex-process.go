package main

import (
	"fmt"
	"os/exec"
	"sort"
	"strconv"
	"strings"
)

func main() {
	ps := exec.Command("ps", "-e", "-opid,ppid,comm")
	output, _ := ps.Output()
	child := make(map[int][]int)
	for i, s := range strings.Split(string(output), "\n") {
		// kill the first line
		if i == 0 {
			continue
		}

		// kill the last line
		if len(s) == 0 {
			continue
		}

		f := strings.Fields(s)
		fpp, _ := strconv.Atoi(f[1]) // parent's id
		fp, _ := strconv.Atoi(f[0])  // child's id
		child[fpp] = append(child[fpp], fp)
	}
	schild := make([]int, len(child))
	i := 0
	for k, _ := range child {
		schild[i] = k
		i++
	}
	sort.Ints(schild)
	for _, ppid := range schild {
		fmt.Printf("Pid %d has %d child", ppid, len(child[ppid]))
		if len(child[ppid]) == 1 {
			fmt.Printf(": %v\n", child[ppid])
			continue
		}
		fmt.Printf("ren: %v\n", child[ppid])
	}
}

/* Output:
Pid 0 has 1 child: [1]
Pid 1 has 189 children: [40 41 43 44 46 48 49 50 53 55 56 60 61 62 63 66 67 69 70 71 73 74 75 76 81 82 85 87 88 89 90 91 92 94 95 96 102 122 124 125 128 139 140 142 145 160 167 168 170 171 172 173 174 175 177 179 182 184 193 200 201 203 205 208 210 211 212 213 214 215 219 220 221 222 223 224 225 226 228 229 231 234 235 237 238 241 242 244 246 249 251 254 255 256 259 260 262 263 265 266 267 268 271 273 274 275 276 277 278 280 281 282 283 285 286 287 288 289 290 293 294 295 296 297 299 301 302 303 305 312 313 314 315 316 317 318 320 327 329 330 340 341 349 350 351 379 405 409 416 1701 2455 2456 2946 3017 5528 8551 13342 13492 13493 17035 17041 18352 18519 18523 18525 19350 19402 19441 19446 21172 21173 21251 21449 21484 21486 21733 25526 25528 25529 25856 25859 25863 25872 42830 47146 56129 56130 58913 26037]
Pid 87 has 2 children: [161 162]
Pid 271 has 1 child: [21458]
Pid 416 has 1 child: [25216]
Pid 19402 has 1 child: [19404]
Pid 21458 has 1 child: [21459]
Pid 21459 has 1 child: [21460]
Pid 25216 has 1 child: [59633]
Pid 25856 has 1 child: [25861]
Pid 59633 has 4 children: [59682 59686 59704 59706]
Pid 59706 has 1 child: [59710]
Pid 59710 has 1 child: [59711]
*/
