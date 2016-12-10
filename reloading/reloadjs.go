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
    const containers = containersForEvent(e.events);
    reloadContainers(containers);
});

function containersForEvent(events) {
    const containers = [];
    document.querySelectorAll("[data-container-reloadon]").forEach(n => {
        // FIXME only processing the first event
        console.log(n, n.dataset.containerReloadon, events[0], n.dataset.containerReloadon === events[0]);
        if (n.dataset.containerReloadon === events[0]) {
            containers.push(n.closest("[data-container-id]"));
        }
    });
    return containers;
}

function reloadContainers(containers) {
    const ids = [];

    containers.forEach(c => {
        const id = c.dataset.containerId
        if (id != null) {
            ids.push(id);
        }
    })

        console.log(containers);
    console.log(ids.join(","));

    if (ids.length > 0) {
        fetchContainers(ids)
            .then(j => {
                console.log(j);
                containers.forEach(c => {
                    id = c.dataset.containerId;
                    console.log("updating "+c+" with:"+j[id]);
                    c.innerHTML = j[id]
                });
            }).catch(e => console.error(e))
    }
}

function fetchContainers(ids) {
    const url = "?c="+ids.join(",");

    const reqData = {
        headers: {
            'Accept': "application/x-container-list"
        }
    };

    return fetch(url, reqData).then(r => r.json())
}

`
