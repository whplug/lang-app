const things = ["1", "2", "3"];

function output(){
    const randomIndex = Math.floor(Math.random() * things.length);
    const random = things[randomIndex];

    document.getElementById('arrayMessage').innerHTML = random;
}
