const app = Vue.createApp({
  setup() {
    const appState = Vue.reactive({
      name: "",
      addr: "",
      token: "",
      canParse: false,
      canParseMulti: false,
      canExport: false,
      priority: 0,
      agentError: "",
      agents: [],
      agentStatusError: "",
    });

    function create() {
      axios
        .post("/api/agent/new", {
          name: appState.name,
          addr: appState.addr,
          token: appState.token,
          can_parse: appState.canParse,
          can_parse_multi: appState.canParseMulti,
          can_export: appState.canExport,
          priority: appState.priority,
        })
        .then(function (response) {
          appState.name = "";
          appState.addr = "";
          appState.token = "";
          appState.canParse = false;
          appState.canParseMulti = false;
          appState.canExport = false;
          appState.priority = 0;
          appState.agentError = "";

          agents();
        })
        .catch(function (error) {
          console.log(error);
          appState.agentError = error.toString();
        });
    }

    function deleteAgent(id) {
      axios
        .post("/api/agent/delete", {
          id: id,
        })
        .then(function (response) {
          agents();
        })
        .catch(function (error) {
          console.log(error);
        });
    }

    function agents() {
      axios
        .post("/api/agent/list", {
          include_status: true,
        })
        .then(function (response) {
          appState.agents = response.data;
          appState.agentStatusError = "";
        })
        .catch(function (error) {
          console.log(error);
          appState.agentStatusError = error.toString();
        });
    }

    Vue.onBeforeMount(() => {
      agents();
    });

    return {
      appState,
      create,
      agents,
      deleteAgent,
    };
  },
});

window.addEventListener("load", async function () {
  app.mount("#app");
});
