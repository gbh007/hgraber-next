const app = Vue.createApp({
  setup() {
    const appState = Vue.reactive({
      name: "",
      addr: "",
      token: "",
      canParse: false,
      canExport: false,
      priority: 0,
      agentError: "",
    });

    function create() {
      axios
        .post("/api/agent/new", {
          name: appState.name,
          addr: appState.addr,
          token: appState.token,
          can_parse: appState.canParse,
          can_export: appState.canExport,
          priority: appState.priority,
        })
        .then(function (response) {
          appState.name = "";
          appState.addr = "";
          appState.token = "";
          appState.canParse = false;
          appState.canExport = false;
          appState.priority = 0;
          appState.agentError = "";
        })
        .catch(function (error) {
          console.log(error);
          appState.agentError = error.toString();
        });
    }

    return {
      appState,
      create,
    };
  },
});

window.addEventListener("load", async function () {
  app.mount("#app");
});
