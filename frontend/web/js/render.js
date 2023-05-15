function renderDiv(inText) {
    const div = document.createElement('div');
    const body = document.body
    div.innerText = inText;
    div.className = 'divClass';
    body.appendChild(div);
}

