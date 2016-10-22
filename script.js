document.addEventListener("click", postAction);

function handleEvents(events) {
    const containers = [];
    document.querySelectorAll("[data-container-id]").forEach(n => {
        console.log(n)
        if (n.children[0] && n.children[0].dataset["events"] === "a") {
            containers.push(n.dataset["containerId"])
        }
    });
    console.log(containers)
    console.log(containers.join(","))
        fetch("?c="+containers.join(","), {
            headers: {
                'Accept': "application/x-container-list"
            }
        })
           .then(r => r.json())
           .then(j => {
               containers.forEach(c => {
                   console.log("updating "+c+" with:"+j[c]);
                   document.querySelector("[data-container-id='"+c+"']").innerHTML = j[c]
               })
           }).catch(e => console.error(e))
}

function postAction(e) {
    console.log(e)
    if (e.target.dataset.containerAction != null) {
        setTimeout(() => handleEvents(["a", "b"]), 100);
    }
}

// click =>
//  request =>
//   respond with events => reload containers
//   respond with error => replace container that triggered event with response
