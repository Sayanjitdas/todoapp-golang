
//function use to delete todos
async function deleteTodo(todoId){
    console.log("TODO ID "+todoId);

    //fetch ajax call
    let response = await fetch("/",{
        method: 'DELETE',
        body: JSON.stringify({"todoId":parseInt(todoId)})
    })
    response = await response.json();
    if(response["Msg"] === "success"){
        document.location.reload();
    }
}