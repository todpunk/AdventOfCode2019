package main

import (
	"bufio"
	"fmt"
	"gopkg.in/alecthomas/kingpin.v2"
	"log"
	"math"
	"os"
	"strconv"
)

var (
	day   = kingpin.Arg("day", "Advent day to run").Int()
)

func main() {
	kingpin.Version("0.0.1")
	kingpin.Parse()
	fmt.Printf("Would run day: %d\n", *day)
	switch *day {
	case 1:
		day1()
	case 2:
		day2()
	case 3:
		day3()
	case 4:
		day4()
	case 5:
		day5()
	default:
		fmt.Println("We don't have that day...")
	}
}

func day1() {
	file, err := os.Open("./day1-input.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	var modules []int64
	var sum int64 = 0
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		var module int64
		module, _ = strconv.ParseInt(scanner.Text(), 10, 32)
		sum += module / 3.0 - 2
		modules = append(modules, module)
	}

	fmt.Printf("Total fuel required: %d\n", sum)
	fmt.Println("But we need to fuel that fuel, so....")

	var newFuel int64
	sum = 0
	for i, _ := range modules {
		for newFuel = modules[i]; newFuel > 0; {
			newFuel = newFuel / 3 - 2
			if newFuel > 0 {
				sum += newFuel
			}
		}
	}
	fmt.Printf("New Total fuel required (calculated per module): %d\n", sum)
}

func getParamValueWithMode(codes []int64, mode int64, location int64) (param int64) {
	switch mode {
	case 0:
		return codes[codes[location]]
	case 1:
		return codes[location]
	default:
		fmt.Printf("Can not get with mode %d as it is unknown at location %d\n", mode, location)
		return -1
	}
}

func setParamWithMode(codes []int64, mode int64, location int64, value int64) (result int64) {
	switch mode {
	case 0:
		codes[codes[location]] = value
	case 1:
		// Yes, I know this should never occur.  But it wasn't called out as explicitly not able to occur.
		codes[location] = value
	default:
		fmt.Printf("Can not set with mode %d as it is unknown at location %d\n", mode, location)
		return -1
	}
	return 0
}

func runIntComp(codes []int64 ) (pos0 int64 ) {
	var i int64

	for i = 0; i < int64(len(codes)-1);  {
		opcode := codes[i] % 100
		mode1 := (codes[i] / 100) % 10
		mode2 := (codes[i] / 1000) % 10
		mode3 := (codes[i] / 10000) % 10
		switch opcode {
		case 99:
			i = int64(len(codes))
		case 1:
			a := getParamValueWithMode(codes, mode1, i+1)
			b := getParamValueWithMode(codes, mode2, i+2)
			result := setParamWithMode(codes, mode3, i+3, a + b)
			if result == -1 {
				i = int64(len(codes))
			}
			i += 4
		case 2:
			a := getParamValueWithMode(codes, mode1, i+1)
			b := getParamValueWithMode(codes, mode2, i+2)
			result := setParamWithMode(codes, mode3, i+3, a * b)
			if result == -1 {
				i = int64(len(codes))
			}
			i += 4
		case 3:
			fmt.Print("Enter number: ")
			var input int64 = 0
			_, _ = fmt.Scanf("%d", &input)
			fmt.Println("Input was:", input)
			result := setParamWithMode(codes, mode1, i+1, input)
			if result == -1 {
				i = int64(len(codes))
			}
			i += 2
		case 4:
			fmt.Printf("%d\n", getParamValueWithMode(codes, mode1, i+1))
			i += 2
		case 5:
			first := getParamValueWithMode(codes, mode1, i+1)
			if first != 0 {
				i = getParamValueWithMode(codes, mode2, i+2)
			} else {
				i += 3
			}
		case 6:
			first := getParamValueWithMode(codes, mode1, i+1)
			if first == 0 {
				i = getParamValueWithMode(codes, mode2, i+2)
			} else {
				i += 3
			}
		case 7:
			first := getParamValueWithMode(codes, mode1, i+1)
			second := getParamValueWithMode(codes, mode2, i+2)
			var result int64
			if first < second {
				result = setParamWithMode(codes, mode3, i+3, 1)
			} else {
				result = setParamWithMode(codes, mode3, i+3, 0)
			}
			i += 4
			if result == -1 {
				i = int64(len(codes))
			}
		case 8:
			first := getParamValueWithMode(codes, mode1, i+1)
			second := getParamValueWithMode(codes, mode2, i+2)
			var result int64
			if first == second {
				result = setParamWithMode(codes, mode3, i+3, 1)
			} else {
				result = setParamWithMode(codes, mode3, i+3, 0)
			}
			i += 4
			if result == -1 {
				i = int64(len(codes))
			}
		default:
			i = int64(len(codes))
			fmt.Printf("This went poorly, opcode: %d at %d\n", opcode, i)
		}
	}
	return codes[0]
}

func day2() {
	rawcodes := []int64{1,0,0,3,1,1,2,3,1,3,4,3,1,5,0,3,2,9,1,19,1,19,5,23,1,23,6,27,2,9,27,31,1,5,31,35,1,35,10,39,1,39,10,43,2,43,9,47,1,6,47,51,2,51,6,55,1,5,55,59,2,59,10,63,1,9,63,67,1,9,67,71,2,71,6,75,1,5,75,79,1,5,79,83,1,9,83,87,2,87,10,91,2,10,91,95,1,95,9,99,2,99,9,103,2,10,103,107,2,9,107,111,1,111,5,115,1,115,2,119,1,119,6,0,99,2,0,14,0}
	var codes = make([]int64, len(rawcodes))
	copy(codes, rawcodes)
	fmt.Println(rawcodes)
	fmt.Println(codes)
	codes[1] = 12
	codes[2] = 2
	fmt.Printf("Position 0: %d\n", runIntComp(codes))

	var pos0, noun, verb int64
	for noun = 1; noun < 100; noun++ {
		for verb = 1; verb < 100; verb++ {
			copy(codes, rawcodes)
			codes[1] = noun
			codes[2] = verb
			pos0 = runIntComp(codes)
			if pos0 == 19690720 {
				fmt.Printf("Noun: %d Verb: %d Result: %d\n", noun, verb, 100 * noun + verb)
				noun = 99
				verb = 99
			}
		}
	}
}

type gridPoint struct{
	x int64
	y int64
}

func Abs(x int64) int64 {
	if x < 0 {
		return -x
	}
	return x
}

func calculateWire(wirePoints map[gridPoint]int64, wire []string){
	var dirs = map[string]gridPoint{
		"R": {1,0},
		"L": {-1,0},
		"U": {0,1},
		"D": {0,-1},
	}

	var current = gridPoint{0,0}
	var totalSteps int64 = 0
	for _, segment := range wire {
		dir := string(segment[0])
		var distance int64
		distance, _ = strconv.ParseInt(segment[1:], 10, 64)
		var i int64
		for i = 0; i < distance; i++ {
			totalSteps += 1
			current.x += dirs[dir].x
			current.y += dirs[dir].y
			wirePoints[current] = totalSteps
		}
	}
}

func calculateOverlap(a, b map[gridPoint]int64) (c map[gridPoint]int64) {
	c = make(map[gridPoint]int64)

	for item, valB := range b {
		if _, ok := a[item]; ok {
			c[item] = a[item] + valB
		}
	}
	return
}

func day3() {
	var wire1 = []string{"R1009","U263","L517","U449","L805","D78","L798","D883","L777","D562","R652","D348","R999","D767","L959","U493","R59","D994","L225","D226","R634","D200","R953","U343","L388","U158","R943","U544","L809","D785","R618","U499","L476","U600","L452","D693","L696","U764","L927","D346","L863","D458","L789","U268","R586","U884","L658","D371","L910","U178","R524","U169","R973","D326","R483","U233","R26","U807","L246","D711","L641","D75","R756","U365","R203","D377","R624","U430","L422","U367","R547","U294","L916","D757","R509","D332","R106","D401","L181","U5","L443","U197","R406","D829","R878","U35","L958","U31","L28","D362","R188","D582","R358","U750","R939","D491","R929","D513","L541","U418","R861","D639","L917","U582","R211","U725","R711","D718","L673","U921","L157","U83","L199","U501","L66","D993","L599","D947","L26","U237","L981","U833","L121","U25","R641","D372","L757","D645","R287","U390","R274","U964","R288","D209","R109","D364","R983","U715","L315","U758","R36","D500","R626","U893","L840","U716","L606","U831","L969","D643","L300","D838","R31","D751","L632","D702","R468","D7","L169","U149","R893","D33","R816","D558","R152","U489","L237","U415","R434","D472","L198","D874","L351","U148","R761","U809","R21","D25","R586","D338","L568","U20","L157","U221","L26","U424","R261","D227","L551","D754","L90","U110","L791","U433","R840","U323","R240","U124","L723","D418","R938","D173","L160","U293","R773","U204","R192","U958","L472","D703","R556","D168","L263","U574","L845","D932","R165","D348","R811","D834","R960","U877","R935","D141","R696","U748","L316","U236","L796","D566","R524","U449","R378","U480","L79","U227","R867","D185","R474","D757","R366","U153","R882","U252","R861","U900","R28","U381","L845","U642","L849","U352","R134","D294","R788","D406","L693","D697","L433","D872","R78","D364","R240","U995","R48","D681","R727","D825","L583","U44","R743","D929","L616","D262","R997","D15","R575","U341","R595","U889","R254","U76","R962","D944","R724","D261","R608","U753","L389","D324","L569","U308","L488","D358","L695","D863","L712","D978","R149","D177","R92"}
	var wire2 = []string{"L1003","D960","L10","D57","R294","U538","R867","D426","L524","D441","R775","U308","R577","D785","R495","U847","R643","D895","R448","U685","L253","U312","L312","U753","L89","U276","R799","D923","L33","U595","R400","U111","L664","D542","R171","U709","L809","D713","L483","U918","L14","U854","L150","D69","L158","D500","L91","D800","R431","D851","L798","U515","L107","U413","L94","U390","L17","U221","L999","D546","L191","U472","L568","U114","L913","D743","L713","D215","L569","D674","L869","U549","L789","U259","L330","D76","R243","D592","L646","U880","L363","U542","L464","D955","L107","U473","R818","D786","R852","U968","R526","D78","L275","U891","R480","U991","L981","D391","R83","U691","R689","D230","L217","D458","R10","U736","L317","D145","R902","D428","R344","U334","R131","D739","R438","D376","L652","U304","L332","D452","R241","D783","R82","D317","R796","U323","R287","D487","L302","D110","R233","U631","R584","U973","L878","D834","L930","U472","R120","U78","R806","D21","L521","U988","R251","D817","R44","D789","R204","D669","R616","D96","R624","D891","L532","U154","R438","U469","R785","D431","R945","U649","R670","D11","R840","D521","L235","D69","L551","D266","L454","U807","L885","U590","L647","U763","R449","U194","R68","U809","L884","U962","L476","D648","L139","U96","L300","U351","L456","D202","R168","D698","R161","U834","L273","U47","L8","D157","L893","D200","L454","U723","R886","U92","R474","U262","L190","U110","L407","D723","R786","D786","L572","D915","L904","U744","L820","D663","R205","U878","R186","U247","L616","D386","R582","U688","L349","D399","R702","U132","L276","U866","R851","D633","R468","D263","R678","D96","L50","U946","R349","D482","R487","U525","R464","U977","L499","D187","R546","U708","L627","D470","R673","D886","L375","U616","L503","U38","L775","D8","L982","D556","R159","U680","L124","U777","L640","D607","R248","D671","L65","D290","R445","U778","L650","U679","L846","D1","L769","U659","R734","D962","R588","U178","R888","D753","R223","U318","L695","D586","R430","D61","R105","U801","R953","U721","L856","U769","R937","D335","R895"}

	var wire1Points = map[gridPoint]int64{}
	var wire2Points = map[gridPoint]int64{}

	calculateWire(wire1Points, wire1)
	calculateWire(wire2Points, wire2)
	var overlaps = calculateOverlap(wire1Points, wire2Points)

	var distance int64 = math.MaxInt64
	var totalSteps int64 = math.MaxInt64
	for overlap, steps := range overlaps {
		var pointDistance = Abs(overlap.x) + Abs(overlap.y)
		if pointDistance < distance {
			distance = pointDistance
		}
		if steps < totalSteps {
			totalSteps = steps
		}
	}
	fmt.Println(distance)
	fmt.Println(totalSteps)
}

func isPotentialPassword(password int64) (correct bool){
	var strPass string = strconv.FormatInt(password, 10)
	var doubleMap = map[int32]int64{}
	var doublesFound int64 = 0

	var last int64 = 0
	for _, c := range strPass {
		var current, _ = strconv.ParseInt(strconv.FormatInt(int64(c), 10), 10, 32)
		if last > current {
			return false
		}
		doubleMap[c] += 1
		// This if statement is only useful for part 2
		if doubleMap[c] == 3 {
			doublesFound--
		}
		if doubleMap[c] == 2 {
			doublesFound++
		}
		last = current
	}
	// If we have made it this far, the only criteria left is if there are doubles somewhere in there
	return doublesFound > 0
}

func day4() {
	var possibility, potentials int64
	potentials = 0

	for possibility = 108457; possibility <= 562041; possibility++ {
		if isPotentialPassword(possibility) {
			potentials++
		}
	}
	fmt.Println(potentials)
}

func day5() {
	rawcodes := []int64{3,225,1,225,6,6,1100,1,238,225,104,0,1101,48,82,225,102,59,84,224,1001,224,-944,224,4,224,102,8,223,223,101,6,224,224,1,223,224,223,1101,92,58,224,101,-150,224,224,4,224,102,8,223,223,1001,224,3,224,1,224,223,223,1102,10,89,224,101,-890,224,224,4,224,1002,223,8,223,1001,224,5,224,1,224,223,223,1101,29,16,225,101,23,110,224,1001,224,-95,224,4,224,102,8,223,223,1001,224,3,224,1,223,224,223,1102,75,72,225,1102,51,8,225,1102,26,16,225,1102,8,49,225,1001,122,64,224,1001,224,-113,224,4,224,102,8,223,223,1001,224,3,224,1,224,223,223,1102,55,72,225,1002,174,28,224,101,-896,224,224,4,224,1002,223,8,223,101,4,224,224,1,224,223,223,1102,57,32,225,2,113,117,224,101,-1326,224,224,4,224,102,8,223,223,101,5,224,224,1,223,224,223,1,148,13,224,101,-120,224,224,4,224,1002,223,8,223,101,7,224,224,1,223,224,223,4,223,99,0,0,0,677,0,0,0,0,0,0,0,0,0,0,0,1105,0,99999,1105,227,247,1105,1,99999,1005,227,99999,1005,0,256,1105,1,99999,1106,227,99999,1106,0,265,1105,1,99999,1006,0,99999,1006,227,274,1105,1,99999,1105,1,280,1105,1,99999,1,225,225,225,1101,294,0,0,105,1,0,1105,1,99999,1106,0,300,1105,1,99999,1,225,225,225,1101,314,0,0,106,0,0,1105,1,99999,8,677,226,224,102,2,223,223,1006,224,329,101,1,223,223,107,677,677,224,1002,223,2,223,1006,224,344,101,1,223,223,8,226,677,224,102,2,223,223,1006,224,359,101,1,223,223,107,226,226,224,102,2,223,223,1005,224,374,1001,223,1,223,1108,677,226,224,1002,223,2,223,1006,224,389,101,1,223,223,107,677,226,224,102,2,223,223,1006,224,404,1001,223,1,223,1107,226,677,224,1002,223,2,223,1006,224,419,1001,223,1,223,108,677,677,224,102,2,223,223,1005,224,434,1001,223,1,223,1008,677,226,224,1002,223,2,223,1006,224,449,1001,223,1,223,7,226,677,224,1002,223,2,223,1006,224,464,1001,223,1,223,1007,677,677,224,102,2,223,223,1005,224,479,1001,223,1,223,1007,226,226,224,1002,223,2,223,1005,224,494,1001,223,1,223,108,226,226,224,1002,223,2,223,1005,224,509,1001,223,1,223,1007,226,677,224,1002,223,2,223,1006,224,524,101,1,223,223,1107,677,677,224,102,2,223,223,1005,224,539,101,1,223,223,1107,677,226,224,102,2,223,223,1005,224,554,1001,223,1,223,108,677,226,224,1002,223,2,223,1006,224,569,1001,223,1,223,1108,226,677,224,1002,223,2,223,1006,224,584,101,1,223,223,8,677,677,224,1002,223,2,223,1006,224,599,1001,223,1,223,1008,226,226,224,102,2,223,223,1006,224,614,101,1,223,223,7,677,677,224,1002,223,2,223,1006,224,629,101,1,223,223,1008,677,677,224,102,2,223,223,1005,224,644,101,1,223,223,7,677,226,224,1002,223,2,223,1005,224,659,101,1,223,223,1108,226,226,224,102,2,223,223,1006,224,674,1001,223,1,223,4,223,99,226}
	var codes = make([]int64, len(rawcodes))
	copy(codes, rawcodes)
	runIntComp(codes)
}
