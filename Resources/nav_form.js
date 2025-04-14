/*ATTENTE DU CHARGEMENT INTEGRALE DE LA PAGE*/
document.addEventListener("DOMContentLoaded", function() { //on ajoute une écoute d'événement au chargement de la page
    console.log("DOM chargé"); //affiche un message console pour vérifie que le contenu est bien chargé

    /*TEST DE PRISE EN CHARGE NAVIGATEUR DU FILTRE FLOU -WEBKIT-*/
    var testEl = document.createElement('div');
    testEl.style.backdropFilter = 'blur(1px)';

        /*si non compatible avec le navigateur, la propriété 'backdrop-filter' reste vide et on applique la classe 'no-backdrop-filter' sans flou*/
        if(testEl.style.backdropFilter === ''){
            document.documentElement.classList.add('no-backdrop-filter');
        }
    
    
    /*ANIMATION DES FORMULAIRES*/
    //sélection des éléments
    const wrapper = document.querySelector('.wrapper');
    const registerLink = document.querySelector('.register-link');
    const loginLink = document.querySelector('.login-link');

    if (registerLink){ //clic sur le lien d'inscription
        registerLink.addEventListener('click', function(e){
            e.preventDefault();
            wrapper.classList.add('active');
        });
    }

    if (loginLink){ //clic sur le lien de connexion
        loginLink.addEventListener('click', function(e){
            e.preventDefault();
            wrapper.classList.add('active');
        }); 
    }
});