'use strict'

module.exports = class Ziraffe {
    constructor() {
        this.lectures = require('./data/lectures').lectures
        this.courses = require('./data/courses').courses
    }

    findLecture = (name, success, failure) => {
        array = this.lectures.reduce((accumulator, current, index) => {
            if(current.name === name) {
                accumulator.push(current)
            }
            return accumulator
        }, [])
        if(array.length == 1) {
            success(accumulator[0])
        } else{
            failure(name)
        }
    }

    similarLectures = (name) => {
        const array = this.lectures.map(item => calculateEditDistance(item.name, name))
        const array1 = array.filter(item => item.distance <= 1)
        if(array1.length == 0){
            return array.filter(item => item.distance <= 2)
        }
        return array1
    }
}

const insert = (node) => 1
const remove = (node) => 1
const update = (string1, string2) => string1 !== string2? 1: 0
let ed = require('edit-distance')

const calculateEditDistance = (name1, name2) => {
    let lev = ed.levenshtein(name1, name2, insert, remove, update)
    return { name1: name1, name2: name2, distance: lev.distance }
}
