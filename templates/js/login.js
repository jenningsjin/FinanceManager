$(document).ready(function(){
	$("#login-form").on("submit", function(e) {
		e.preventDefault();
		$.ajax({
			type: "POST",
			url: "/login/",
			data: $(this).serialize(),
			success: function(data) {
				console.log("success");
			},
			error: function(data) {
				console.log("error");
				$("#errorMessage").text(data.responseText);
			}
		});
	});
});