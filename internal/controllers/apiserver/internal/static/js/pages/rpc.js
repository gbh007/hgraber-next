const app = Vue.createApp({
  setup() {
    const appState = Vue.reactive({
      deduplicateFiles: null,
      deduplicateFilesError: "",
      removeDetachedFiles: null,
      removeDetachedFilesError: "",
      removeMismatchFiles: null,
      removeMismatchFilesError: "",

      systemInfoData: null,
      systemInfoError: "",
    });

    function deduplicateFiles() {
      axios
        .post("/api/system/rpc/deduplicate/files")
        .then(function (response) {
          appState.deduplicateFiles = response.data;
          appState.deduplicateFilesError = "";
        })
        .catch(function (error) {
          console.log(error);
          appState.deduplicateFilesError = error.toString();
        });
    }

    function removeDetachedFiles() {
      axios
        .post("/api/system/rpc/remove/detached-files")
        .then(function (response) {
          appState.removeDetachedFiles = response.data;
          appState.removeDetachedFilesError = "";
        })
        .catch(function (error) {
          console.log(error);
          appState.removeDetachedFilesError = error.toString();
        });
    }

    function removeMismatchFiles() {
      axios
        .post("/api/system/rpc/remove/mismatch-files")
        .then(function (response) {
          appState.removeMismatchFiles = response.data;
          appState.removeMismatchFilesError = "";
        })
        .catch(function (error) {
          console.log(error);
          appState.removeMismatchFilesError = error.toString();
        });
    }

    async function systemInfo() {
      axios
        .get("/api/system/info")
        .then(function (response) {
          let data = response.data;

          appState.systemInfoData = data;
          appState.systemInfoError = "";
        })
        .catch(function (error) {
          console.log(error);
          appState.systemInfoError = error.toString();
        });
    }

    async function setRunnerCount() {
      console.log(appState.systemInfoData.monitor.workers);

      axios
        .post("/api/system/worker/config", {
          runners_count: appState.systemInfoData.monitor.workers.map((e) => ({
            name: e.name,
            count: e.runners,
          })),
        })
        .then(function (response) {
          appState.systemInfoError = "";
        })
        .catch(function (error) {
          console.log(error);
          appState.systemInfoError = error.toString();
        });
    }
    Vue.onBeforeMount(() => {
      systemInfo();
    });

    return {
      appState,
      deduplicateFiles,
      removeDetachedFiles,
      removeMismatchFiles,
      systemInfo,
      setRunnerCount,
    };
  },
});

window.addEventListener("load", async function () {
  app.mount("#app");
});
