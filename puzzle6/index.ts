import * as fs from 'fs';


const school: number[] = [0, 0, 0, 0, 0, 0, 0, 0, 0]

// read input
const fileContent = fs.readFileSync('input.txt')

let count = 0
// set up the init school of fish
fileContent.toString().split(',').forEach(age => {
    const a = parseInt(age)
    school[a]++
    count++
})
console.log(count)
console.log(school)

// simulate school growth
let newBirths: number = 0
for (let day = 0; day < 256; day++) {
    // shift
    newBirths = 0

    for (let i = 0; i < school.length; i++) {
        if (i == 0) {
            newBirths = school[i]
        } else {
            school[i - 1] = school[i]
        }
    }
    school[8] = newBirths
    school[6] += newBirths
    count += newBirths
    // console.log(count, school)
}

console.log("Number of fish: ", count)