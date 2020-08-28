var counter : number = 0;

setInterval(()=>{
	console.log("firing " + counter);
	counter++;
}, Math.random() * 1000);

function drawBoard(board:number[][]) {
	console.log("board", JSON.stringify(board));
}