lucide.createIcons();

document.addEventListener("alpine:init", () => {
  Alpine.data("music", () => ({
    current: 0,
    lyrics: [],

    metadata: {
      title: "",
      artist: "",
      album: "",
      language: "",
    },

    updateTime() {
      this.current = this.$el.currentTime;
    },

    update() {
      const audio = this.$refs["music-player"];
      const file = this.$el.files?.[0];

      if (file && audio) {
        audio.src = URL.createObjectURL(file);
      }
    },

    addLyric() {
      const formData = new FormData(this.$el);

      this.$refs["lyric-input"].value = "";

      this.lyrics.unshift({
        value: formData.get("value"),
        time: formatSeconds(this.current),
      });
    },

    removeLyric(index) {
      this.lyrics.splice(index, 1);
    },

    updateLyric(index, property) {
      this.lyrics[index][property] = this.$el.value;
    },

    async _export() {
      const resp = await fetch("/lrc", {
        method: "POST",
        body: JSON.stringify({ lyrics: this.lyrics }),
      });
      downloadFile(await resp.blob(), this.metadata.title);
    },
  }));
});

function downloadFile(blob, title) {
  const url = URL.createObjectURL(blob);

  const el = document.createElement("a");
  el.href = url;
  el.download = `${title} Lyric.lrc`;
  el.click();
  el.remove();
}

function formatSeconds(seconds) {
  const minutes = Math.floor(seconds / 60)
    .toString()
    .padStart(2, "0");
  var seconds = (seconds % 60).toFixed(2).padStart(5, "0");

  return `${minutes}:${seconds}`;
}
