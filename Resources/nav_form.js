/*GESTION DES PAGES FORUM ET LOG_REG*/

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
    //sélection des éléments du DOM
    const authLayout = document.querySelector('.auth-layout');
    
    if(authLayout){ //on vérifie que l'on est bien sur la page d'authentification
        const registerLink = document.querySelector('.register-link');
        const loginLink = document.querySelector('.login-link');    
    
        if (registerLink){ //clic sur le lien d'inscription pour basculer entre les formulaires
            registerLink.addEventListener('click', function(e){
                e.preventDefault();
                authLayout.classList.add('active');

                //active les animations pour le formulaire d'inscription
                const registerAnimations = document.querySelectorAll('.form-box.register .animation');
                registerAnimations.forEach(function(el, index){
                    el.style.opacity = '0';
                    setTimeout(function(){
                        el.style.opacity = '1';
                    },500 + (index * 100)); //délais après la transition
                });
            });
        }

        if (loginLink){ //clic sur le lien de connexion
            loginLink.addEventListener('click', function(e){
                e.preventDefault();
                authLayout.classList.remove('active');

                //active les animation pour le formulaire de connexion
                const loginAnimations = document.querySelectorAll('.form-box.login .animation');
                loginAnimations.forEach(function(el, index){
                    el.style.opacity ='0';
                    setTimeout(function(){
                        el.style.opacity = '1';
                    }, 500 + (index * 100));
                });
            }); 
        }

        //active les animations au chargement de la page
        const animations= document.querySelectorAll('.animation');
        animations.forEach(function(el, index){
            el.style.opacity = '0';
            //ajout d'un délais basé sur l'index en cas d'incompatibilté CSS
            setTimeout(function(){
                el.style.opacity = '1';
            }, 100 + (index * 100));
        });
    }
});