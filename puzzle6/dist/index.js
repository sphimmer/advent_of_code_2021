"use strict";
var __createBinding = (this && this.__createBinding) || (Object.create ? (function(o, m, k, k2) {
    if (k2 === undefined) k2 = k;
    Object.defineProperty(o, k2, { enumerable: true, get: function() { return m[k]; } });
}) : (function(o, m, k, k2) {
    if (k2 === undefined) k2 = k;
    o[k2] = m[k];
}));
var __setModuleDefault = (this && this.__setModuleDefault) || (Object.create ? (function(o, v) {
    Object.defineProperty(o, "default", { enumerable: true, value: v });
}) : function(o, v) {
    o["default"] = v;
});
var __importStar = (this && this.__importStar) || function (mod) {
    if (mod && mod.__esModule) return mod;
    var result = {};
    if (mod != null) for (var k in mod) if (k !== "default" && Object.prototype.hasOwnProperty.call(mod, k)) __createBinding(result, mod, k);
    __setModuleDefault(result, mod);
    return result;
};
Object.defineProperty(exports, "__esModule", { value: true });
const fs = __importStar(require("fs"));
const school = [0, 0, 0, 0, 0, 0, 0, 0, 0];
// read input
const fileContent = fs.readFileSync('input.txt');
let count = 0;
// set up the init school of fish
fileContent.toString().split(',').forEach(age => {
    const a = parseInt(age);
    school[a]++;
    count++;
});
console.log(count);
console.log(school);
// simulate school growth
let newBirths = 0;
for (let day = 0; day < 256; day++) {
    // shift
    newBirths = 0;
    for (let i = 0; i < school.length; i++) {
        if (i == 0) {
            newBirths = school[i];
        }
        else {
            school[i - 1] = school[i];
        }
    }
    school[8] = newBirths;
    school[6] += newBirths;
    count += newBirths;
    // console.log(count, school)
}
console.log("Number of fish: ", count);
//# sourceMappingURL=index.js.map