package reloading

var ReloadScript = `
//////////////////////////////////////////
// Framework code

// Event input
function postEvent(name) {
    const e = new Event("appEvent");
    e.events = [name];
    console.log("posting:"+name);
    document.dispatchEvent(e);
}

// Event -> Reload containers mapped to event
document.addEventListener("appEvent", e => {
    console.log("handling", e.events);

    const containers = [];
    document.querySelectorAll("[data-container-id]").forEach(n => {
        containers.push(n)
    });
    reloadContainers(e.events, containers);
});


function reloadContainers(tagNames, containers) {
    fetchContainers(tagNames)
        .then(j => {
            console.log(j);
            containers.forEach(c => {
                id = c.dataset.containerId;
                console.log("updating "+c+" with:"+j[id]);
                if(j[id] == null) {
                    return
                }
                c.innerHTML = j[id]
            });
        }).catch(e => console.error(e))
}

function fetchContainers(tagNames) {
    const url = "?containersByTags="+tagNames.join(",");

    const reqData = {
        headers: {
            'Accept': "application/x-container-list"
        }
    };

    return fetch(url, reqData).then(r => r.json())
}

`
