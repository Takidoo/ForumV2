document.addEventListener("DOMContentLoaded", () => {
  const form = document.getElementById("threadform");

  form.addEventListener("submit", async (e) => {
    e.preventDefault(); 

    const formData = new FormData(form);
    const urlEncoded = new URLSearchParams(formData);

    try {
      const response = await fetch("/api/CreateThread", {
        method: "POST",
        headers: {
          "Content-Type": "application/x-www-form-urlencoded"
        },
        body: urlEncoded,
      });

      if (response.ok) {
        window.location.href = "/";
      } else {
        const res = await response.text();
        alert("Erreur : " + res);
      }
    } catch (err) {
      console.error("Erreur réseau :", err);
      alert("Erreur réseau.");
    }
  });
});
