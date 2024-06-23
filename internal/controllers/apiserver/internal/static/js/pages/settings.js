const app = Vue.createApp({
  setup() {
    const appState = Vue.reactive({
      settings: JSON.parse(localStorage.getItem("settings")) || {},
      bookOnPage: 0,
    });

    Vue.onBeforeMount(() => {
      appState.bookOnPage = appState.settings.book_on_page || 12;
    });

    function saveSettings() {
      appState.settings.book_on_page = appState.bookOnPage;
      localStorage.setItem("settings", JSON.stringify(appState.settings));
    }

    return {
      appState,
      saveSettings,
    };
  },
});

window.addEventListener("load", async function () {
  app.mount("#app");
});
