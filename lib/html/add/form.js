

var userPhones = []


function isEmpty(obj) {
	if (typeof obj == "undefined" || obj == "undefined" || obj == null || obj == "") {
		return true;
	}
	return false;
}

function getUser() {
	var dup = {};

	$.post(
		"/rep/api/v1/op?type=user&op=range",
		{ "type": "user", "op": "range" },
		function (data, status) {
			JSON.parse(data).forEach(item => {
				var phone = String(item["phone"]);
				var name = String(item["name"]);

				if (!isEmpty(phone) && !dup[k]) {
					var k = phone;
					if (!isEmpty(name)) {
						k = k + ":" + name
					}
					userPhones.push({ title: k });
					dup[k] = true;
				}
			});
		}
	)
}


var formFieldValues = new Array();
var formField = ["type", "colour", "opendirection", "size", "shape", "glass", "priceunit", "priceset"]

function getForm() {
	var dup = {}

	$.post(
		"/rep/api/v1/op?type=form&op=range",
		{ "type": "form", "op": "range" },
		function (data, status) {
			JSON.parse(data).forEach(item => {
				formField.forEach(field => {
					if (!dup.hasOwnProperty(field)) {
						dup[field] = {}
					}

					var k = String(item[field]);
					if (!isEmpty(k) && !dup[field][k]) {
						dup[field][k] = true;
						if (!formFieldValues.hasOwnProperty(field)) {
							formFieldValues[field] = [];
						}
						formFieldValues[field].push({ title: k });
					}
				});
			});
		}
	)
}

$.ajaxSettings.async = false;
// 获取全量用户信息
getUser()
// 获取表单信息
getForm()
$.ajaxSettings.async = true;



// 客户名单
$('#f-phone').bigAutocomplete({ data: userPhones });
$('#f-type').bigAutocomplete({ data: formFieldValues["type"] });
$('#f-colour').bigAutocomplete({ data: formFieldValues["colour"] });
$('#f-opendirection').bigAutocomplete({ data: formFieldValues["opendirection"] });
$('#f-size').bigAutocomplete({ data: formFieldValues["size"] });
$('#f-shape').bigAutocomplete({ data: formFieldValues["shape"] });
$('#f-glass').bigAutocomplete({ data: formFieldValues["glass"] });
$('#f-priceunit').bigAutocomplete({ data: formFieldValues["priceunit"] });
$('#f-priceset').bigAutocomplete({ data: formFieldValues["priceset"] });

function checkPhone(e) {
	e.value = e.value.split(":")[0];
}

function addForm() {
	var form_addform = document.getElementById('form-addform');

	var newElement = document.createElement("input")
	newElement.setAttribute("type", "hidden")
	newElement.name = "type"
	newElement.value = "form"
	form_addform.appendChild(newElement)

	var newElement = document.createElement("input")
	newElement.setAttribute("type", "hidden")
	newElement.name = "op"
	newElement.value = "put"
	form_addform.appendChild(newElement)

	form_addform.submit()
}