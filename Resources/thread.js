document.addEventListener('DOMContentLoaded', () => {

document.getElementById("sidebar-reply-send").addEventListener("click", function () {
const content = document.getElementById("sidebar-reply-textarea").value;

const params = new URLSearchParams();
params.append("message", content);
const url = new URLSearchParams(window.location.search);
params.append("thread_id", url.get("thread_id"));

fetch("/api/CreatePost", {
method: "POST",
headers: {
    "Content-Type": "application/x-www-form-urlencoded"
},
body: params.toString()
})
.then(response => {
if (response.ok) {
    location.reload();
} else {
    alert("Erreur lors de l'envoi.");
}
})
.catch(error => {
console.error("Erreur réseau :", error);
alert("Erreur réseau.");
});
});


document.getElementById("like-button").addEventListener("click", function () {
    
    const params = new URLSearchParams();
    const url = new URLSearchParams(window.location.search);
    params.append("thread_id", url.get("thread_id"));
    
    fetch("/api/LikeThread", {
    method: "POST",
    headers: {
        "Content-Type": "application/x-www-form-urlencoded"
    },
    body: params.toString()
    })
    .then(response => {
    if (response.ok) {
        location.reload();
    } else {
        alert("Vous avez déjà liker ce thread");
    }
    })
    .catch(error => {
    console.error("Erreur réseau :", error);
    alert("Erreur réseau.");
    });
    });
});