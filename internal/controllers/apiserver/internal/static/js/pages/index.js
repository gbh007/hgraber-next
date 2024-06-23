const app = Vue.createApp({
  setup() {
    const appState = Vue.reactive({
      state: {},
      stateError: "",
      urlsRaw: "",
      urlsResult: {},
      urlsError: "",
      stateIntervalID: null,
    });

    async function remakeInfo() {
      axios
        .get("/api/system/info")
        .then(function (response) {
          let data = response.data;

          appState.state = data;
          appState.stateError = "";
        })
        .catch(function (error) {
          console.log(error);
          appState.stateError = error.toString();
        });
    }

    async function loadBooks() {
      axios
        .post("/api/system/handle", {
          urls: appState.urlsRaw.split("\n").map((s) => s.trim()),
        })
        .then(function (response) {
          let data = response.data;
          appState.urlsResult = data;
          appState.urlsError = "";
          appState.urlsRaw = (data.not_handled || []).join("\n");
        })
        .catch(function (error) {
          console.log(error);
          appState.urlsError = error.toString();
        });
    }

    Vue.onMounted(() => {
      remakeInfo();
      appState.stateIntervalID = setInterval(remakeInfo, 10000);
    });

    Vue.onUnmounted(() => {
      clearInterval(appState.stateIntervalID);
    });

    return {
      appState,
      remakeInfo,
      loadBooks,
    };
  },
});

window.addEventListener("load", async function () {
  app.mount("#app");
});
