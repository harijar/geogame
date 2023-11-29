import {AuthRequest} from "./api";

async function onTelegramAuth(user) {
    try {
        await AuthRequest(user);
    } catch (error) {
        alert(error);
        return
    }
    alert('Logged in as ' + user.first_name + ' ' + user.last_name + ' (' + user.id + (user.username ? ', @' + user.username : '') + ')');
}