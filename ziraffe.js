'use strict'

const createArray = (name1, name2) => {
    const array = new Array(name1.length + 1)
    for(let i = 0; i <= name1.length; i++){
        array[i] = new Array(name2.length + 1).fill(0)
        array[i][0] = i
    }
    for(let j = 0; j <= name2.length; j++) array[0][j] = j
    return array
}

const findDefaultFunc = (orig, defaultFunc) => {
    return orig === undefined? defaultFunc: orig
}

const restRequirements = (course, credits) => {
    return course.requirements.filter(item => credits.indexOf(item) < 0)
}
const gotRequirements = (course, credits) => {
    return course.requirements.filter(item => credits.indexOf(item) > 0)
}

module.exports = class Ziraffe {
    constructor() {
        this.lectures = require('./data/lectures').lectures
        this.courses = require('./data/courses').courses
    }

    gotCredits(course, credits) {
        let sum = 0
        credits.forEach(name => {
            if (course.requirements.indexOf(name) >= 0) {
                this.findLecture(name, (lecture) => {
                    sum = sum + lecture.credit
                })
            }
        })
        return sum
    }

    checkEachCourse(course, credits) {
        return {
            name: course.name,
            requirements: course.requirements,
            got_credits: this.gotCredits(course, credits),
            got_requirements: gotRequirements(course, credits),
            rest_requirements: restRequirements(course, credits),
        }
    }

    graduationCheck(credits) {
        const available = credits.filter(name => this.isFindLecture(name))
        return this.courses.map(course => this.checkEachCourse(course, available))
    }

    findLecture(name, success, failure) {
        const array = this.lectures.reduce((accumulator, current) => {
            if (current.name === name) {
                accumulator.push(current)
            }
            return accumulator
        }, [])
        if (array.length == 1) {
            success(array[0])
        } else {
            failure(name)
        }
    }

    isFindLecture(name) {
        const array = this.lectures.reduce((accumulator, current) => {
            if (current.name === name) {
                accumulator.push(current)
            }
            return accumulator
        }, [])
        return array.length == 1
    }

    similarLectures(name) {
        const array = this.lectures.map(item => this.calculateEditDistance(item.name, name))
        const array1 = array.filter(item => item.distance <= 1)
        if (array1.length == 0) {
            return array.filter(item => item.distance <= 2)
        }
        return array1
    }

    calculateEditDistance(name1, name2) {
        const distance = this.levenshtein(name1, name2)
        return { name1: name1, name2: name2, distance: distance }
    }

    levenshtein(name1, name2, insertFunc, removeFunc, updateFunc) {
        const array = createArray(name1, name2)
        insertFunc = findDefaultFunc(insertFunc, (node) => 1)
        removeFunc = findDefaultFunc(removeFunc, (node) => 1)
        updateFunc = findDefaultFunc(updateFunc, (string1, string2) => string1 !== string2? 1: 0)
        for(let i = 1; i <= name1.length; i++) {
            for(let j = 1; j <= name2.length; j++) {
                const d1 = array[i - 1][j] + removeFunc(name1.charAt(i - 1))
                const d2 = array[i][j - 1] + insertFunc(name2.charAt(j - 1))
                const d3 = updateFunc(name1.charAt(i - 1), name2.charAt(j - 1))
                array[i][j] = Math.min(d1, d2, d3)
            }
        }
        for(let i = 0; i <= name1.length; i++) {
            let line = ""
            for(let j = 0; j <= name2.length; j++) {
                let item = "    " + array[i][j]
                line = line + item.slice(-3)
            }
            console.log(line)
        }
        return array[name1.length][name2.length]
    }

}
