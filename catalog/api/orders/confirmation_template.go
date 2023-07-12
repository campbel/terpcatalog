package orders

import (
	"text/template"

	"github.com/campbel/terpcatalog/shared/types"
)

func generateConfirmationTemplate(order types.Order) (string, error) {
	var items []struct {
		Name     string
		Quantity int
		Price    float64
	}
	for _, item := range order.Items {
		items = append(items, struct {
			Name     string
			Quantity int
			Price    float64
		}{
			Name:     item.Strain.Name,
			Quantity: item.Quantity,
			Price:    item.Strain.Price,
		})
	}

	var total float64
	for _, item := range items {
		total += item.Price * float64(item.Quantity)
	}

	return renderTemplate(parsedConfirmationTemplate, map[string]any{
		"Items": items,
		"Total": total,
	})
}

func Cost(price float64, quantity int) float64 {
	return price * float64(quantity)
}

var parsedConfirmationTemplate = template.Must(template.New("email").Funcs(template.FuncMap{"cost": Cost}).Parse(confirmationTemplate))

var confirmationTemplate = `
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
		</style>
	</head>
	<body>
		<div class="container">
			<div class="header">
				<h1>TerpScout</h1>
			</div>
			<div class="body">
				<h2>Order Confirmation</h2>
				<p>Thank you for your order! We will send you an email when your order is ready for delivery.</p>
				<h3>Order Details</h3>
				<table>
					<tr>
						<th>Item</th>
						<th>Quantity</th>
						<th>Price</th>
						<th>Total</th>
					</tr>
					{{ range .Items }}
					<tr>
						<td>{{ .Name }}</td>
						<td>{{ .Quantity }}</td>
						<td>{{ .Price }}</td>
						<td>{{ cost .Price .Quantity }}</td>
					</tr>
					{{ end }}
				</table>
				<h3>Order Total: {{ .Total }}</h3>
			</div>
			<div class="footer">
				<p>TerpScout</p>
			</div>
		</div>
	</body>
</html>
`
