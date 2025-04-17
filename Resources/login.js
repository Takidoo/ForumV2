const wrapper = document.querySelector('.wrapper')
const registerLink = document.querySelector('.register-link')
const loginLink = document.querySelector('.login-link')

registerLink.onclick = () => {
    wrapper.classList.add('active')
}

loginLink.onclick = () => {
    wrapper.classList.remove('active')
}

const formlogin = document.getElementById("form-login");

  formlogin.addEventListener("submit", async (e) => {
    e.preventDefault(); 

    const formData = new FormData(formlogin);
    const urlEncoded = new URLSearchParams(formData);

    try {
      const response = await fetch("/api/login", {
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

  const formregister = document.getElementById("form-register");

  formregister.addEventListener("submit", async (e) => {
    e.preventDefault(); 

    const formData = new FormData(formregister);
    const urlEncoded = new URLSearchParams(formData);

    try {
      const response = await fetch("/api/register", {
        method: "POST",
        headers: {
          "Content-Type": "application/x-www-form-urlencoded"
        },
        body: urlEncoded,
      });

      if (response.ok) {
        alert("Compte créer avec succès")
        window.location.href = "/login";
      } else {
        const res = await response.text();
        alert("Erreur : " + res);
      }
    } catch (err) {
      console.error("Erreur réseau :", err);
      alert("Erreur réseau.");
    }
  });