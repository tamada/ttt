'use strict'

const insert = (node) => 1
const remove = (node) => 1
const update = (string1, string2) => string1 !== string2 ? 1 : 0
let ed = require('edit-distance')

const calculateEditDistance = (name1, name2) => {
    let lev = ed.levenshtein(name1, name2, insert, remove, update)
    return { name1: name1, name2: name2, distance: lev.distance }
}

const gotCredits = (course, credits) => {
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

const restRequirements = (course, credits) => {
    return course.requirements.filter(item => credits.indexOf(item) < 0)
}
const gotRequirements = (course, credits) => {
    return course.requirements.filter(item => credits.indexOf(item) > 0)
}

const checkCourse = (course, credits) => {
    return {
        name: course.name,
        requirements: course.requirements,
        got_credits: gotCredits(course, credits),
        got_requirements: gotRequirements(course, credits),
        rest_requirements: restRequirements(course, credits),
    }
}

module.exports = class Ziraffe {
    constructor() {
        this.lectures = require('./data/lectures').lectures
        this.courses = require('./data/courses').courses
    }


    graduationCheck(credits) {
        const available = credits.filter(name => this.isFindLecture(name))
        return this.courses.map(course => checkCourse(course, available))
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
        const array = this.lectures.map(item => calculateEditDistance(item.name, name))
        const array1 = array.filter(item => item.distance <= 1)
        if (array1.length == 0) {
            return array.filter(item => item.distance <= 2)
        }
        return array1
    }
}
