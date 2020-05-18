// Function to change the link and title content once they are clicked
function setLink(){
	// Get the url of the page
	var path = window.location.pathname;
	// Get the title and anchor elements
	var title = document.getElementById("pastedition_title");
	var link = document.getElementById("pastedition_link");
	// Get the content of each element
	var title_content = title.textContent;
	// Check the content for the title tag
	if (path.includes("2019")) {
		// Change title and link to a the 2020 edition
		title.textContent = " 2020 Track Overview ";
		link.textContent = "For 2020 overview, click here.";
		link.href = "index.html";
		//window.alert("First_loop");
	} else if (path.includes("index")) {
		// Change title and link to a the 2019 edition
		title.textContent = " 2019 Track Overview ";
		link.textContent = "For 2019 overview, click here.";
		link.href = "2019.html";
	}
}
