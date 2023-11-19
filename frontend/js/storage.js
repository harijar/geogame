export function GetPrompts() {
    let prompts = localStorage.getItem('prompts');
    if (prompts) {
        const decoded = JSON.parse(prompts);
        if (decoded) {
            return decoded;
        }
    }
    return [];
}

export function SavePrompts(prompts) {
    localStorage.setItem('prompts', JSON.stringify(prompts));
}