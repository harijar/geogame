import {AuthRequest} from "./api.js";

window.auth = auth

async function auth(user) {
    try {
        await AuthRequest(user);
    } catch (error) {
        alert(error);
        return
    }
    alert('Logged in as ' + user.first_name + ' ' + user.last_name + ' (' + user.id + (user.username ? ', @' + user.username : '') + ')');
}