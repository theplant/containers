//////////////////////////////////////////
// Application code

// dummy method for triggering some kind of "action"
document.addEventListener("click", postAction);

function postAction(e) {
    console.log(e)
    const event = e.target.dataset.containerEvent
    if (event != null) {
        setTimeout(() => postEvent(event), 100);
    }
}
