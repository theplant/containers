# Containers

# Reloading Spec

Containers can be reloaded when "events" occur. Triggering events is application-specific, it could come from a server action, or from JS.

Demo works by:

1. Elements (buttons, links, etc) define a `data-container-event` attribute, value is  an event name.

2. Elements inside a container define a `data-container-reloadon` attribute, value is an event name.

3. An event listener is attached to the document that listens for `appEvent` DOM events.

4. When clicking on an element defined in step 1, "app" code calls `postEvent` with the event name specified in `data-container-event`. This dispatches a DOM event of `appEvent` with the passed event as a parameter.

5. The listener from step 3 receives the `appEvent`, extracts the event's name, and finds all elements with `[data-container-reloadon=<event name>]`. For each element, it finds the nearest ancestor "container" (element with a `data-container-id`).

6. The list of containers found in step 5 is passed to `reloadContainers` which requests only the specific containers from the server.

7. The JSON result of 6 is used to replace (via `innerHTML`) the existing content of the contaners found in step 5 with the data from the server in step 6.


