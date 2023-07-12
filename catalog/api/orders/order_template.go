package orders

import (
	"text/template"

	"github.com/campbel/terpcatalog/shared/types"
)

func generateOrderTemplate(order types.Order) (string, error) {
	return renderTemplate(parsedOrderTemplate, order)
}

var parsedOrderTemplate = template.Must(template.New("email").Parse(orderTemplate))

var orderTemplate = `
<html>
	<head>
		<style>
			* {
				font-family: sans-serif;
			}
			.container {
				max-width: 600px;
				margin: 0 auto;
			}
			.header {
				background-color: #f5f5f5;
				padding: 20px;
			}
			.header h1 {
				margin: 0;
			}
			.body {
				padding: 20px;
			}
			.footer {
				background-color: #f5f5f5;
 				padding: 20px;
 			}
 			.footer p {
				margin: 0;
			}
		</style>
	</head>
	<body>
		<div class="container">
			<div class="header">
				<h1>Order Submitted</h1>
			</div>
			<div class="body">
			<h3>Information:</h3>
			<ul>
				<li>Company: {{.Information.Company}}</li>
				<li>License Number: {{.Information.LicenseNumber}}</li>
				<li>Email: {{.Information.Email}}</li>
				<li>Phone: {{.Information.Phone}}</li>
				<li>Address: {{.Information.Address.Street}}</li>
				<li>City: {{.Information.Address.City}}</li>
				<li>State: {{.Information.Address.State}}</li>
				<li>Postal: {{.Information.Address.Postal}}</li>
			</ul>
			<h3>Items</h3>
			<table>
				<thead>
					<tr>
						<th>Producer</th>
						<th>ID</th>
						<th>Name</th>
						<th>Quantity</th>
						<th>Price</th>
					</tr>
				</thead>
				<tbody>
					{{range .Items}}
					<tr>
						<td>{{.Strain.ProducerID}}</td>
						<td>{{.Strain.ID}}</td>
						<td>{{.Strain.Name}}</td>
						<td>{{.Quantity}}</td>
						<td>{{.Strain.Price}}</td>
					</tr>
					{{end}}
				</tbody>
			</table>
			</div>
			<div class="footer">
				<p>Get to it!</p>
			</div>
		</div>
	</body>
</html>
`
