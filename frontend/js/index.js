import {Auth, CheckAuth} from "./auth.js"

window.onload = async function() {
    await CheckAuth();
    document.getElementById("userProfileButton").onclick = async function() {
        window.location.href = '/user/';
    }
    document.getElementById("playButton").onclick = async function() {
        window.location.href = 'game.html';
    }
}
async function auth(user) {
    await Auth(user);
}