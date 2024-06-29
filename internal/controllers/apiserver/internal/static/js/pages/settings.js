const app = Vue.createApp({
  setup() {
    const appState = Vue.reactive({
      settings: JSON.parse(localStorage.getItem("settings")) || {},
      bookOnPage: 0,
      token: "",
      tokenError: "",
    });

    Vue.onBeforeMount(() => {
      appState.bookOnPage = appState.settings.book_on_page || 12;
    });

    function saveSettings() {
      appState.settings.book_on_page = appState.bookOnPage;
      localStorage.setItem("settings", JSON.stringify(appState.settings));
    }

    function login() {
      axios
        .post("/api/user/login", { token: appState.token })
        .then(function (response) {
          appState.token = "";
          appState.tokenError = "";
        })
        .catch(function (error) {
          console.log(error);
          appState.tokenError = error.toString();
        });
    }

    return {
      appState,
      saveSettings,
      login,
    };
  },
});

window.addEventListener("load", async function () {
  app.mount("#app");
});
