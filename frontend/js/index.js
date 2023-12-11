import {Auth, CheckAuth} from "./auth.js"

window.onload = async function() {await CheckAuth()}
async function auth(user) {
    await Auth(user);
}