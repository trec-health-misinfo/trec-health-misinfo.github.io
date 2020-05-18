// Function to change the link and title content once they are clicked
document.getElementById("pastedition_link").onclick = function(){
	// Get the title and anchor elements
	var title = document.getElementById("pastedition_title");
	var link = document.getElementById("pastedition_link");
	// Get the content of each element
	var title_content = title.textContent;
	// Check the content for the title tag
	if (title_content == " 2019 Track Overview ") {
		// Change title and link to a the 2020 edition
		title.textContent = " 2020 Track Overview ";
		link.textContent = "For 2020 overview, click here.";
		link.href = "index.html";
		break;
	} else if (title_content == " 2020 Track Overview ") {
		// Change title and link to a the 2019 edition
		title.textContent = " 2019 Track Overview ";
		link.textContent = "For 2019 overview, click here.";
		link.href = "2019.html";
		break;
	}
}
