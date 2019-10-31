'use strict'

const ziraffe = new Ziraffe()

const insertItem = (item) => {
    $('#lectures-list').append(`<input type="check" value="false">${item.name}</input>（配当学年: ${item.grade}，単位数: ${item.credit}）`)
}

const initialize = () => {
    $('#lectures-list').empty()
    ziraffe.lectures.sort((item1, item2) => {
        if (item1.grade < item2.grade) return -1
        else if (item1.grade > item2.grade) return 1
        return 0
    }).forEach(item => insertItem(item))
}

(() => {
  initialize()
)
