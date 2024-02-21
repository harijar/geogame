import {UsersRequest} from "../api.js"

const userElement = document.createElement("div")
userElement.appendChild(document.createElement("p"))
userElement.appendChild(document.createElement("p"))
userElement.children[0].classList.add("userNickname")
for (let i = 0; i < 3; i++) {
    userElement.appendChild(document.createElement("br"))
}

window.onload = async function() {
    document.getElementById("mainMenuButton").onclick = async function() {
        window.location.href = '/index.html'
    }
    let pageNumber = 0
    await getUsers(pageNumber)
    document.getElementById("prevPage").onclick = async function() {
        pageNumber--
        await getUsers(pageNumber)
    }
    document.getElementById("nextPage").onclick = async function() {
        pageNumber++
        await getUsers(pageNumber)
    }
}

async function getUsers(pageNumber) {
    if (pageNumber === 0) {
        document.getElementById("prevPage").disabled = true
    }
    let data = await UsersRequest(pageNumber)
    let users = data['users']
    if (users.length === 0) {
        document.getElementById("nextPage").disabled = true
        pageNumber--
        return
    }

    for (let i = 0; i < users.length; i++) {
        let user = users[i]
        let el = userElement
        el.children[0].appendChild(document.createTextNode(user['nickname']))
        el.children[1].appendChild(document.createTextNode(user['last_seen']))
        document.getElementById("users").appendChild(el)
    }
}