// Copyright 2019 The KubeSphere Authors. All rights reserved.
// Use of this source code is governed by a Apache license
// that can be found in the LICENSE file.

package constants

const (
	EmailKubeSphereNotifyActiveTemplateZh = `
	<!DOCTYPE html>
	<html>
		<head>
			<meta name="viewport" content="width=device-width" />
			<meta http-equiv="Content-Type" content="text/html; charset=UTF-8" />
			<title>Simple Transactional Email</title>
			<style>
				/* -------------------------------------
					GLOBAL RESETS
				------------------------------------- */
				/*All the styling goes here*/
				img {
					border: none;
					-ms-interpolation-mode: bicubic;
					max-width: 100%;
				}
	
				body {
					background-color: #ecf1f7;
					font-family: Roboto, PingFang SC, Lantinghei SC, Helvetica Neue,
						Helvetica, Arial, Microsoft YaHei, 微软雅黑, STHeitiSC-Light, simsun,
						宋体, WenQuanYi Zen Hei, WenQuanYi Micro Hei, sans-serif;
					-webkit-font-smoothing: antialiased;
					font-size: 14px;
					line-height: 1.4;
					margin: 0;
					padding: 0;
					-ms-text-size-adjust: 100%;
					-webkit-text-size-adjust: 100%;
					color: #242e42;
				}
	
				table {
					border-collapse: separate;
					mso-table-lspace: 0pt;
					mso-table-rspace: 0pt;
					width: 100%;
				}
				table td {
					font-size: 14px;
					vertical-align: top;
				}
	
				/* -------------------------------------
					BODY & CONTAINER
				------------------------------------- */
	
				.body {
					background-color: #eff0f5;
					width: 100%;
				}
	
				/* Set a max-width, and make it display as block so it will automatically stretch to that width, but will also shrink down on a phone or something */
				.container {
					display: block;
					margin: 0 auto !important;
					/* makes it centered */
					max-width: 780px;
					padding: 10px;
					padding-top: 80px;
					width: 780px;
				}
	
				/* This should also be a block element, so that it will fill 100% of the .container */
				.content {
					box-sizing: border-box;
					display: block;
					margin: 0 auto;
					max-width: 780px;
					padding: 10px;
				}
	
				/* -------------------------------------
					HEADER, FOOTER, MAIN
				------------------------------------- */
				.main {
					background: #ffffff;
					border-radius: 4px;
					border: solid 1px #e3e9ef;
					background-color: #ffffff;
				}
	
				.wrapper {
					box-sizing: border-box;
					padding: 48px;
				}
	
				.content-block {
					padding-bottom: 10px;
					padding-top: 10px;
				}
	
				.footer {
					clear: both;
					margin-top: 14px;
					text-align: center;
					width: 100%;
				}
				.footer td,
				.footer p,
				.footer span,
				.footer a {
					color: #8c96ad;
					font-size: 12px;
					text-align: center;
				}
				.gray {
					color: #8c96ad;
				}
				.logo {
					display: inline-block;
					vertical-align: middle;
				}
	
				/* -------------------------------------
					TYPOGRAPHY
				------------------------------------- */
				h1,
				h2,
				h3,
				h4 {
					color: #000000;
					font-weight: 400;
					line-height: 1.4;
					margin: 0;
					margin-bottom: 30px;
				}
	
				h1 {
					font-size: 35px;
					font-weight: 300;
					text-align: center;
					text-transform: capitalize;
				}
	
				p,
				ul,
				ol {
					font-size: 14px;
					font-weight: normal;
					line-height: 2;
					margin: 0;
					margin-bottom: 15px;
				}
				p li,
				ul li,
				ol li {
					list-style-position: inside;
					margin-left: 5px;
				}
	
				a {
					color: #329dce;
					font-weight: bold;
					text-decoration: none;
				}
	
				/* -------------------------------------
					BUTTONS
				------------------------------------- */
				.btn {
					box-sizing: border-box;
					width: 100%;
				}
				.btn > tbody > tr > td {
					padding-bottom: 15px;
				}
				.btn table {
					width: auto;
				}
				.btn table td {
					background-color: #ffffff;
					border-radius: 5px;
					text-align: center;
				}
				.btn a {
					background-color: #ffffff;
					border: solid 1px #3498db;
					border-radius: 5px;
					box-sizing: border-box;
					color: #3498db;
					cursor: pointer;
					display: inline-block;
					font-size: 14px;
					font-weight: bold;
					margin: 0;
					padding: 12px 25px;
					text-decoration: none;
					text-transform: capitalize;
				}
	
				.btn-primary table td {
					background-color: #3498db;
				}
	
				.btn-primary a {
					background-color: #3498db;
					border-color: #3498db;
					color: #ffffff;
				}
	
				/* -------------------------------------
					OTHER STYLES THAT MIGHT BE USEFUL
				------------------------------------- */
				.last {
					margin-bottom: 0;
				}
	
				.first {
					margin-top: 0;
				}
	
				.align-center {
					text-align: center;
				}
	
				.align-right {
					text-align: right;
				}
	
				.align-left {
					text-align: left;
				}
	
				.clear {
					clear: both;
				}
	
				.mt0 {
					margin-top: 0;
				}
	
				.mb0 {
					margin-bottom: 0;
				}
	
				.preheader {
					color: transparent;
					display: none;
					height: 0;
					max-height: 0;
					max-width: 0;
					opacity: 0;
					overflow: hidden;
					mso-hide: all;
					visibility: hidden;
					width: 0;
				}
	
				.powered-by a {
					text-decoration: none;
				}
	
				hr {
					border: 0;
					border-bottom: 1px solid #eff0f5;
					margin: 50px 0 12px;
				}
				.linkBtn {
					border-radius: 2px;
					box-shadow: 0 1px 3px 0 rgba(73, 33, 173, 0.16),
						0 1px 2px 0 rgba(52, 57, 69, 0.03);
					background-color: #242e42;
					color: #fff;
					line-height: 20px;
					padding: 8px 20px;
				}
				.link {
					font-size: 12px;
					font-style: normal;
					font-stretch: normal;
					line-height: 28px;
					letter-spacing: normal;
				}
				.platform {
					font-size: 14px;
					font-weight: 500;
					font-style: normal;
					font-stretch: normal;
					line-height: 20px;
					letter-spacing: normal;
					color: #343945;
					margin-left: 12px;
				}
				.line1 {
					margin-top: 22px;
					margin-bottom: 16px;
				}
				.line2 {
					line-height: 2;
					margin-top: 16px;
					margin-bottom: 20px;
				}
				.line3 {
					margin-bottom: 40px;
					margin-top: 16px;
				}
				.line4 {
					margin-bottom: 0px;
				}
				.line5 {
					margin-top: 0px;
				}
				.line6 {
					margin-bottom: 0px;
				}
	
				/* -------------------------------------
					RESPONSIVE AND MOBILE FRIENDLY STYLES
				------------------------------------- */
				@media only screen and (max-width: 620px) {
					table[class='body'] h1 {
						font-size: 28px !important;
						margin-bottom: 10px !important;
					}
					table[class='body'] p,
					table[class='body'] ul,
					table[class='body'] ol,
					table[class='body'] td,
					table[class='body'] span,
					table[class='body'] a {
						font-size: 16px !important;
					}
					table[class='body'] .wrapper,
					table[class='body'] .article {
						padding: 10px !important;
					}
					table[class='body'] .content {
						padding: 0 !important;
					}
					table[class='body'] .container {
						padding: 0 !important;
						width: 100% !important;
					}
					table[class='body'] .main {
						border-left-width: 0 !important;
						border-radius: 0 !important;
						border-right-width: 0 !important;
					}
					table[class='body'] .btn table {
						width: 100% !important;
					}
					table[class='body'] .btn a {
						width: 100% !important;
					}
					table[class='body'] .img-responsive {
						height: auto !important;
						max-width: 100% !important;
						width: auto !important;
					}
				}
	
				/* -------------------------------------
					PRESERVE THESE STYLES IN THE HEAD
				------------------------------------- */
				@media all {
					.ExternalClass {
						width: 100%;
					}
					.ExternalClass,
					.ExternalClass p,
					.ExternalClass span,
					.ExternalClass font,
					.ExternalClass td,
					.ExternalClass div {
						line-height: 100%;
					}
					.apple-link a {
						color: inherit !important;
						font-family: inherit !important;
						font-size: inherit !important;
						font-weight: inherit !important;
						line-height: inherit !important;
						text-decoration: none !important;
					}
					.btn-primary table td:hover {
						background-color: #34495e !important;
					}
					.btn-primary a:hover {
						background-color: #34495e !important;
						border-color: #34495e !important;
					}
				}
			</style>
		</head>
		<body class="">
			<span class="preheader"
				>This is preheader text. Some clients will show this text as a
				preview.</span
			>
			<table
				role="presentation"
				border="0"
				cellpadding="0"
				cellspacing="0"
				class="body"
			>
				<tr>
					<td>&nbsp;</td>
					<td class="container">
						<div class="content">
							<!-- START CENTERED WHITE CONTAINER -->
							<table role="presentation" class="main">
								<!-- START MAIN CONTENT AREA -->
								<tr>
									<td class="wrapper">
										<table
											role="presentation"
											border="0"
											cellpadding="0"
											cellspacing="0"
										>
											<tr>
												<td>
													<p styles="line-height: 30px;">
														<img
															width="150"
															height="30"
															src="data:image/svg+xml;base64,PHN2ZyB3aWR0aD0iNDI4IiBoZWlnaHQ9IjkwIiB4bWxucz0iaHR0cDovL3d3dy53My5vcmcvMjAwMC9zdmciPjxnIGZpbGw9Im5vbmUiIGZpbGwtcnVsZT0iZXZlbm9kZCI+PGcgZmlsbD0iIzQ0M0Q0RSI+PHBhdGggZD0iTTMwNy41MzQgMjNoLTMuNjY2Yy0uOCAwLTEuNDY4LjY4NC0xLjQ2OCAxLjUwNFY0Mi4wMUgyOTAuNlYyNC41MDRjMC0uODItLjYtMS41MDQtMS40NjYtMS41MDRoLTMuNjY4Yy0uOCAwLTEuNDY3LjY4NC0xLjQ2NyAxLjUwNHY0MS45OWMwIC44MjMuNjY3IDEuNTA2IDEuNDY3IDEuNTA2aDMuNjY4Yy44NjYgMCAxLjQ2Ni0uNjgzIDEuNDY2LTEuNTA1VjQ4Ljc4SDMwMi40djE3LjcxNWMwIC44MjIuNjY3IDEuNTA1IDEuNDY4IDEuNTA1aDMuNjY2Yy44NjcgMCAxLjQ2Ni0uNjgzIDEuNDY2LTEuNTA1di00MS45OWMwLS44Mi0uNi0xLjUwNS0xLjQ2Ni0xLjUwNU0yMTggNjIuNzM0YzAtLjgyMS0uNjc3LTEuNTA1LTEuNDg5LTEuNTA1aC0xNi44MVY0OC44ODZoMTYuODFjLjgxMiAwIDEuNDg4LS42MTUgMS40ODgtMS41MDV2LTMuNzZjMC0uODIyLS42NzctMS41MDYtMS40ODktMS41MDZoLTE2LjgwOVYyOS43N2gxNi44MWMuODEgMCAxLjQ4OC0uNjE1IDEuNDg4LTEuNTA0di0zLjc2YzAtLjgyMy0uNjc3LTEuNTA2LTEuNDg5LTEuNTA2aC0xOC4yOTUtMy43MjVjLS44MTQgMC0xLjQ5LjY4My0xLjQ5IDEuNTA0djQxLjk5YzAgLjgyMi42NzYgMS41MDYgMS40OSAxLjUwNkgyMTYuNTFjLjgxMiAwIDEuNDg5LS42MTYgMS40ODktMS41MDV2LTMuNzYxek0zNDAgNjIuNzM0YzAtLjgyMS0uNjc3LTEuNTA1LTEuNDg5LTEuNTA1aC0xNi44MVY0OC44ODZoMTYuODFjLjgxIDAgMS40ODgtLjYxNSAxLjQ4OC0xLjUwNXYtMy43NmMwLS44MjItLjY3Ny0xLjUwNi0xLjQ4OS0xLjUwNmgtMTYuODA5VjI5Ljc3aDE2LjgxYy44MSAwIDEuNDg4LS42MTUgMS40ODgtMS41MDR2LTMuNzZjMC0uODIzLS42NzctMS41MDYtMS40ODktMS41MDZoLTE4LjI5NS0zLjcyNWMtLjgxNCAwLTEuNDkuNjgzLTEuNDkgMS41MDR2NDEuOTljMCAuODIyLjY3NiAxLjUwNiAxLjQ5IDEuNTA2SDMzOC41MWMuODEyIDAgMS40ODktLjYxNiAxLjQ4OS0xLjUwNXYtMy43NjF6TTQwNCA2Mi43MzRjMC0uODIxLS42NzctMS41MDUtMS40ODktMS41MDVoLTE2LjgxVjQ4Ljg4NmgxNi44MWMuODEyIDAgMS40ODgtLjYxNSAxLjQ4OC0xLjUwNXYtMy43NmMwLS44MjItLjY3Ny0xLjUwNi0xLjQ4OS0xLjUwNmgtMTYuODA5VjI5Ljc3aDE2LjgxYy44MSAwIDEuNDg4LS42MTUgMS40ODgtMS41MDR2LTMuNzZjMC0uODIzLS42NzctMS41MDYtMS40ODktMS41MDZoLTE4LjI5NS0zLjcyNWMtLjgxMyAwLTEuNDkuNjgzLTEuNDkgMS41MDR2NDEuOTljMCAuODIyLjY3NyAxLjUwNiAxLjQ5IDEuNTA2SDQwMi41MWMuODEyIDAgMS40ODktLjYxNiAxLjQ4OS0xLjUwNXYtMy43NjF6TTIyOS40NjcgMzQuNzcydjIuNTg2YzAgMS45NzQgMS44NjcgMy4xMyA0IDQuMDE1bDguODY2IDMuODFjMy4zMzMgMS40OTggNS42NjcgNC4xNSA1LjY2NyA4LjIzNHY0LjgzQzI0OCA2Mi4yNjQgMjM5LjczMiA2OSAyMzUuMzMyIDY5Yy0yLjk5OSAwLTcuNTMyLTIuMTc4LTExLjA2NC00LjE1MS0uODAxLS40MDgtMS4zMzQtMS4zNjItLjg2Ny0yLjQ1bDEuMTM0LTIuMzE0Yy40NjYtLjk1MiAxLjQ2NS0xLjE1NyAyLjMzMi0uNjggMi44NjYgMS40OTcgNi42NjYgMy4zMzQgOC40NjUgMy4zMzQgMiAwIDYuMjAxLTMuNDcgNi4yMDEtNS40NDR2LTIuOTkzYzAtMi4zMTQtMS43MzMtMy42MDctNC4xMzMtNC41NTlsLTguMjY3LTMuNTRjLTMuMi0xLjM2LTYuMTMzLTQuMTUtNi4xMzMtNy45NlYzMy44MmMwLTQuNDI0IDguNzMzLTEwLjgyIDEyLjkzNC0xMC44MiAyLjggMCA3LjMzMiAyLjExIDEwLjMzMyAzLjc0My45MzMuNDc3IDEuMiAxLjU2NS43OTggMi4zODFMMjQ2IDMxLjQzOGMtLjMzMy44MTYtMS40IDEuMDktMi4zMzIuNjgtMi4yNjctMS4wODgtNi4wNjctMi44NTctNy43MzMtMi44NTctMS45MzQgMC02LjQ2NyAzLjI2Ni02LjQ2NyA1LjUxMU0zNTkuNDUyIDI4Ljg4MmgtNi42NTd2MTUuMDQ2aDYuNzI1YzIuODE1IDAgNi4xMDktNC45OTIgNi4xMDktNy42NiAwLTIuNTk4LTMuNDMyLTcuMzg2LTYuMTc3LTcuMzg2bS0uMjc1IDIwLjkyN2gtNi4zODJ2MTYuNjg4YzAgLjg4OC0uNjE4IDEuNTAzLTEuNTggMS41MDNoLTMuNjM2Yy0uODI0IDAtMS41NzktLjYxNS0xLjU3OS0xLjUwM3YtNDAuMzVjMC0xLjU3MyAxLjg1NC0zLjE0NyAzLjM2My0zLjE0N2gxMC45ODJDMzY1LjM1NCAyMyAzNzIgMzAuNzg1IDM3MiAzNS45MTRjMCAzLjQyLTMuMDA4IDguNTYtNi4yMzMgMTAuNTQ0IDIuNjA3IDEuMzY3IDYuMTc2IDQuOTkyIDYuMTc2IDguMzQ0djExLjY5NWExLjUyIDEuNTIgMCAwIDEtMS41MSAxLjUwM2gtMy43NzRhMS41MiAxLjUyIDAgMCAxLTEuNTEtMS41MDN2LTEwLjY3YzAtMi4zOTQtMy4zNjItNi4wMTgtNS45NzItNi4wMThNMjY3LjA4OCAyOC44ODJoLTYuNDc3djE1LjA0Nmg2LjU0M2MyLjczOSAwIDUuNTYtNC45OTIgNS41Ni03LjY2IDAtMi41OTgtMi45NTYtNy4zODYtNS42MjYtNy4zODZtLS4yNjcgMjAuOTI3aC02LjIxdjE2LjY4OGMwIC44ODgtLjYwMiAxLjUwMy0xLjUzNiAxLjUwM2gtMy41NGMtLjggMC0xLjUzNS0uNjE1LTEuNTM1LTEuNTAzdi00MC4zNWMwLTEuNTczIDEuODAzLTMuMTQ3IDMuMjcyLTMuMTQ3aDEwLjY4NEMyNzIuODMgMjMgMjc5IDMxLjAyOCAyNzkgMzYuMTU2YzAgMy40Mi0yLjQ3MyA4LjAyLTUuMzMgMTAuNjgtMS40MDUgMS4zMDgtMy42MzYgMi45NzMtNi44NDkgMi45NzNNMTQyLjQ4NCA2OUMxMzguMTM2IDY5IDEzMCA2My4yMzcgMTMwIDU4LjI2M1YyNC44ODRjMC0uODI4LjcyMi0xLjUxNCAxLjUxNy0xLjUxNGgzLjQ4NWMuODUyIDAgMS41MS42ODYgMS41MSAxLjUxNFY1Ny4zN2MwIDIuNDE4IDMuOTMgNS4zNDkgNS45NzIgNS4zNDkgMi4wNDIgMCA1Ljk5OC0zLjMwMiA1Ljk5OC01LjcyVjI0LjUxNGMwLS44My41OTItMS41MTQgMS4zODgtMS41MTRoMy42MTNjLjc5NSAwIDEuNTE3LjY4NSAxLjUxNyAxLjUxNHYzMy4zOEMxNTUgNjIuNzE3IDE0Ni44MjcgNjkgMTQyLjQ4NCA2OU0xNzQuMTU3IDYyLjEyaC02LjM5NFY0Ny4wNzJoNi40NjJjMi44MDEgMCA2LjA4MSA0LjczIDYuMDgxIDcuMzk2IDAgMi42LTMuNDE2IDcuNjUtNi4xNDkgNy42NXptLTYuMzk0LTMzLjI0aDYuMzk0YzIuNzMzIDAgNi4xNDkgMy45MTggNi4xNDkgNi41MTggMCAyLjY2Ny0zLjI4IDcuMDIyLTYuMDgxIDcuMDIyaC02LjQ2MlYyOC44OHptMTIuOTg4IDE1Ljg2N2MzLjExMy0yLjA1MiA2LjI0OS02LjAzOSA2LjI0OS05LjM1QzE4NyAzMC4yNyAxODAuMDMyIDIzIDE3NS4wNDYgMjNoLTEwLjdjLTEuNTAxIDAtMy4zNDYgMS41NzMtMy4zNDYgMy4xNDV2NDAuMzUyYzAgLjg5Ljc1MSAxLjUwMyAxLjU3MSAxLjUwM2gxMi43MDZDMTgwLjI2MyA2OCAxODcgNTkuNiAxODcgNTQuNDdjMC0zLjMxLTMuMTM2LTcuNjQ3LTYuMjQ5LTkuNzIzek0xMDguOTA5IDQ1Ljc1bDE0LjcxNy0xNy4yNDRjLjUzLS42MjIuNTEyLTEuNTQtLjE1My0yLjEyMmwtMi44MTQtMi40NThhMS40ODggMS40ODggMCAwIDAtMi4wOTguMTU1bC0xMi44NyAxNS4wNzhWMjQuNTA1YzAtLjgyLS42MDgtMS41MDUtMS40ODYtMS41MDVoLTMuNzE4Yy0uODExIDAtMS40ODcuNjg0LTEuNDg3IDEuNTA1djQxLjk5MmMwIC44Mi42NzYgMS41MDMgMS40ODcgMS41MDNoMy43MThjLjg3OCAwIDEuNDg3LS42ODMgMS40ODctMS41MDNWNTIuMzM5bDEyLjg2OSAxNS4wOGExLjQ5IDEuNDkgMCAwIDAgMi4wOTguMTU0bDIuODE0LTIuNDU4Yy42NjUtLjU4Mi42ODMtMS41MDEuMTUzLTIuMTIxbC0xNC43MTctMTcuMjQ1ek00MjAuNTEgMjYuNDI1aC0xLjg0djIuNDM0aDEuODRjLjU5NCAwIDEuMjQ1LS40ODEgMS4yNDUtMS4xOSAwLS43NjMtLjY1LTEuMjQ0LTEuMjQ1LTEuMjQ0em0xLjEzMSA2LjAyOGwtMS43ODItMi43MTdoLTEuMTl2Mi43MTdoLS45NjF2LTYuODc3aDIuODAyYzEuMTYgMCAyLjIzNS44MiAyLjIzNSAyLjA5NCAwIDEuNTI5LTEuMzU4IDIuMDM4LTEuNzU0IDIuMDM4bDEuODQgMi43NDVoLTEuMTl6TTQyMCAyMy45MDZBNS4wNTYgNS4wNTYgMCAwIDAgNDE0LjkwNiAyOSA1LjA5IDUuMDkgMCAwIDAgNDIwIDM0LjA5NGE1LjA5IDUuMDkgMCAwIDAgNS4wOTQtNS4wOTNBNS4wNTYgNS4wNTYgMCAwIDAgNDIwIDIzLjkwNnpNNDIwIDM1Yy0zLjMxMSAwLTYtMi42ODktNi02IDAtMy4zMzkgMi42ODktNiA2LTYgMy4zNCAwIDYgMi42NjEgNiA2IDAgMy4zMTEtMi42NiA2LTYgNnoiLz48L2c+PHBhdGggZmlsbD0iIzAwQTk3MSIgZD0iTTY1IDcxLjY0N2wtMTktMTF2MjJ6TTY1IDE5LjY0N2wtMTktMTF2MjJ6TTE5LjY3OCA0NS42NDdMMzcgMzUuNTU2VjMuNjQ3TDEgMjQuNjJ2NDIuMDU1bDM2IDIwLjk3MlY1NS43Mzl6Ii8+PHBhdGggZmlsbD0iIzAwQTk3MSIgZD0iTTM3IDQ1LjY0N2wzNyAyMXYtNDJ6Ii8+PC9nPjwvc3ZnPg=="
															alt="KubeSphere®️"
														/>
													</p>
													<strong>告警消息</strong>
													<p class="line1">
														资源 {{.ResourceName}} {{.RuleName}} 告警
														{{.CumulatedCount}}次
													</p>
													<p class="line2">首次发生时间：{{.FirstTime}}</p>
													<p>最后一次发生时间：{{.LastTime}}</p>
													<p>最后一次实时值：{{.LastValue}}</p>
													<hr />
													<p class="gray">
														* 此为系统邮件请勿回复
													</p>
												</td>
											</tr>
										</table>
									</td>
								</tr>
							</table>
	
							<div class="footer">
								<table
									role="presentation"
									border="0"
									cellpadding="0"
									cellspacing="0"
								>
									<tr>
										<td class="content-block">
											<span class="apple-link"
												>KubeSphere®️ 2019 All Rights Reserved.</span
											>
										</td>
									</tr>
								</table>
							</div>
						</div>
					</td>
					<td>&nbsp;</td>
				</tr>
			</table>
		</body>
	</html>	
`

	EmailKubeSphereNotifyResumeTemplateZh = `
<!DOCTYPE html>
	<html>
		<head>
			<meta name="viewport" content="width=device-width" />
			<meta http-equiv="Content-Type" content="text/html; charset=UTF-8" />
			<title>Simple Transactional Email</title>
			<style>
				/* -------------------------------------
					GLOBAL RESETS
				------------------------------------- */
				/*All the styling goes here*/
				img {
					border: none;
					-ms-interpolation-mode: bicubic;
					max-width: 100%;
				}
	
				body {
					background-color: #ecf1f7;
					font-family: Roboto, PingFang SC, Lantinghei SC, Helvetica Neue,
						Helvetica, Arial, Microsoft YaHei, 微软雅黑, STHeitiSC-Light, simsun,
						宋体, WenQuanYi Zen Hei, WenQuanYi Micro Hei, sans-serif;
					-webkit-font-smoothing: antialiased;
					font-size: 14px;
					line-height: 1.4;
					margin: 0;
					padding: 0;
					-ms-text-size-adjust: 100%;
					-webkit-text-size-adjust: 100%;
					color: #242e42;
				}
	
				table {
					border-collapse: separate;
					mso-table-lspace: 0pt;
					mso-table-rspace: 0pt;
					width: 100%;
				}
				table td {
					font-size: 14px;
					vertical-align: top;
				}
	
				/* -------------------------------------
					BODY & CONTAINER
				------------------------------------- */
	
				.body {
					background-color: #eff0f5;
					width: 100%;
				}
	
				/* Set a max-width, and make it display as block so it will automatically stretch to that width, but will also shrink down on a phone or something */
				.container {
					display: block;
					margin: 0 auto !important;
					/* makes it centered */
					max-width: 780px;
					padding: 10px;
					padding-top: 80px;
					width: 780px;
				}
	
				/* This should also be a block element, so that it will fill 100% of the .container */
				.content {
					box-sizing: border-box;
					display: block;
					margin: 0 auto;
					max-width: 780px;
					padding: 10px;
				}
	
				/* -------------------------------------
					HEADER, FOOTER, MAIN
				------------------------------------- */
				.main {
					background: #ffffff;
					border-radius: 4px;
					border: solid 1px #e3e9ef;
					background-color: #ffffff;
				}
	
				.wrapper {
					box-sizing: border-box;
					padding: 48px;
				}
	
				.content-block {
					padding-bottom: 10px;
					padding-top: 10px;
				}
	
				.footer {
					clear: both;
					margin-top: 14px;
					text-align: center;
					width: 100%;
				}
				.footer td,
				.footer p,
				.footer span,
				.footer a {
					color: #8c96ad;
					font-size: 12px;
					text-align: center;
				}
				.gray {
					color: #8c96ad;
				}
				.logo {
					display: inline-block;
					vertical-align: middle;
				}
	
				/* -------------------------------------
					TYPOGRAPHY
				------------------------------------- */
				h1,
				h2,
				h3,
				h4 {
					color: #000000;
					font-weight: 400;
					line-height: 1.4;
					margin: 0;
					margin-bottom: 30px;
				}
	
				h1 {
					font-size: 35px;
					font-weight: 300;
					text-align: center;
					text-transform: capitalize;
				}
	
				p,
				ul,
				ol {
					font-size: 14px;
					font-weight: normal;
					line-height: 2;
					margin: 0;
					margin-bottom: 15px;
				}
				p li,
				ul li,
				ol li {
					list-style-position: inside;
					margin-left: 5px;
				}
	
				a {
					color: #329dce;
					font-weight: bold;
					text-decoration: none;
				}
	
				/* -------------------------------------
					BUTTONS
				------------------------------------- */
				.btn {
					box-sizing: border-box;
					width: 100%;
				}
				.btn > tbody > tr > td {
					padding-bottom: 15px;
				}
				.btn table {
					width: auto;
				}
				.btn table td {
					background-color: #ffffff;
					border-radius: 5px;
					text-align: center;
				}
				.btn a {
					background-color: #ffffff;
					border: solid 1px #3498db;
					border-radius: 5px;
					box-sizing: border-box;
					color: #3498db;
					cursor: pointer;
					display: inline-block;
					font-size: 14px;
					font-weight: bold;
					margin: 0;
					padding: 12px 25px;
					text-decoration: none;
					text-transform: capitalize;
				}
	
				.btn-primary table td {
					background-color: #3498db;
				}
	
				.btn-primary a {
					background-color: #3498db;
					border-color: #3498db;
					color: #ffffff;
				}
	
				/* -------------------------------------
					OTHER STYLES THAT MIGHT BE USEFUL
				------------------------------------- */
				.last {
					margin-bottom: 0;
				}
	
				.first {
					margin-top: 0;
				}
	
				.align-center {
					text-align: center;
				}
	
				.align-right {
					text-align: right;
				}
	
				.align-left {
					text-align: left;
				}
	
				.clear {
					clear: both;
				}
	
				.mt0 {
					margin-top: 0;
				}
	
				.mb0 {
					margin-bottom: 0;
				}
	
				.preheader {
					color: transparent;
					display: none;
					height: 0;
					max-height: 0;
					max-width: 0;
					opacity: 0;
					overflow: hidden;
					mso-hide: all;
					visibility: hidden;
					width: 0;
				}
	
				.powered-by a {
					text-decoration: none;
				}
	
				hr {
					border: 0;
					border-bottom: 1px solid #eff0f5;
					margin: 50px 0 12px;
				}
				.linkBtn {
					border-radius: 2px;
					box-shadow: 0 1px 3px 0 rgba(73, 33, 173, 0.16),
						0 1px 2px 0 rgba(52, 57, 69, 0.03);
					background-color: #242e42;
					color: #fff;
					line-height: 20px;
					padding: 8px 20px;
				}
				.link {
					font-size: 12px;
					font-style: normal;
					font-stretch: normal;
					line-height: 28px;
					letter-spacing: normal;
				}
				.platform {
					font-size: 14px;
					font-weight: 500;
					font-style: normal;
					font-stretch: normal;
					line-height: 20px;
					letter-spacing: normal;
					color: #343945;
					margin-left: 12px;
				}
				.line1 {
					margin-top: 22px;
					margin-bottom: 16px;
				}
				.line2 {
					line-height: 2;
					margin-top: 16px;
					margin-bottom: 20px;
				}
				.line3 {
					margin-bottom: 40px;
					margin-top: 16px;
				}
				.line4 {
					margin-bottom: 0px;
				}
				.line5 {
					margin-top: 0px;
				}
				.line6 {
					margin-bottom: 0px;
				}
	
				/* -------------------------------------
					RESPONSIVE AND MOBILE FRIENDLY STYLES
				------------------------------------- */
				@media only screen and (max-width: 620px) {
					table[class='body'] h1 {
						font-size: 28px !important;
						margin-bottom: 10px !important;
					}
					table[class='body'] p,
					table[class='body'] ul,
					table[class='body'] ol,
					table[class='body'] td,
					table[class='body'] span,
					table[class='body'] a {
						font-size: 16px !important;
					}
					table[class='body'] .wrapper,
					table[class='body'] .article {
						padding: 10px !important;
					}
					table[class='body'] .content {
						padding: 0 !important;
					}
					table[class='body'] .container {
						padding: 0 !important;
						width: 100% !important;
					}
					table[class='body'] .main {
						border-left-width: 0 !important;
						border-radius: 0 !important;
						border-right-width: 0 !important;
					}
					table[class='body'] .btn table {
						width: 100% !important;
					}
					table[class='body'] .btn a {
						width: 100% !important;
					}
					table[class='body'] .img-responsive {
						height: auto !important;
						max-width: 100% !important;
						width: auto !important;
					}
				}
	
				/* -------------------------------------
					PRESERVE THESE STYLES IN THE HEAD
				------------------------------------- */
				@media all {
					.ExternalClass {
						width: 100%;
					}
					.ExternalClass,
					.ExternalClass p,
					.ExternalClass span,
					.ExternalClass font,
					.ExternalClass td,
					.ExternalClass div {
						line-height: 100%;
					}
					.apple-link a {
						color: inherit !important;
						font-family: inherit !important;
						font-size: inherit !important;
						font-weight: inherit !important;
						line-height: inherit !important;
						text-decoration: none !important;
					}
					.btn-primary table td:hover {
						background-color: #34495e !important;
					}
					.btn-primary a:hover {
						background-color: #34495e !important;
						border-color: #34495e !important;
					}
				}
			</style>
		</head>
		<body class="">
			<span class="preheader"
				>This is preheader text. Some clients will show this text as a
				preview.</span
			>
			<table
				role="presentation"
				border="0"
				cellpadding="0"
				cellspacing="0"
				class="body"
			>
				<tr>
					<td>&nbsp;</td>
					<td class="container">
						<div class="content">
							<!-- START CENTERED WHITE CONTAINER -->
							<table role="presentation" class="main">
								<!-- START MAIN CONTENT AREA -->
								<tr>
									<td class="wrapper">
										<table
											role="presentation"
											border="0"
											cellpadding="0"
											cellspacing="0"
										>
											<tr>
												<td>
													<p styles="line-height: 30px;">
														<img
															width="150"
															height="30"
															src="data:image/svg+xml;base64,PHN2ZyB3aWR0aD0iNDI4IiBoZWlnaHQ9IjkwIiB4bWxucz0iaHR0cDovL3d3dy53My5vcmcvMjAwMC9zdmciPjxnIGZpbGw9Im5vbmUiIGZpbGwtcnVsZT0iZXZlbm9kZCI+PGcgZmlsbD0iIzQ0M0Q0RSI+PHBhdGggZD0iTTMwNy41MzQgMjNoLTMuNjY2Yy0uOCAwLTEuNDY4LjY4NC0xLjQ2OCAxLjUwNFY0Mi4wMUgyOTAuNlYyNC41MDRjMC0uODItLjYtMS41MDQtMS40NjYtMS41MDRoLTMuNjY4Yy0uOCAwLTEuNDY3LjY4NC0xLjQ2NyAxLjUwNHY0MS45OWMwIC44MjMuNjY3IDEuNTA2IDEuNDY3IDEuNTA2aDMuNjY4Yy44NjYgMCAxLjQ2Ni0uNjgzIDEuNDY2LTEuNTA1VjQ4Ljc4SDMwMi40djE3LjcxNWMwIC44MjIuNjY3IDEuNTA1IDEuNDY4IDEuNTA1aDMuNjY2Yy44NjcgMCAxLjQ2Ni0uNjgzIDEuNDY2LTEuNTA1di00MS45OWMwLS44Mi0uNi0xLjUwNS0xLjQ2Ni0xLjUwNU0yMTggNjIuNzM0YzAtLjgyMS0uNjc3LTEuNTA1LTEuNDg5LTEuNTA1aC0xNi44MVY0OC44ODZoMTYuODFjLjgxMiAwIDEuNDg4LS42MTUgMS40ODgtMS41MDV2LTMuNzZjMC0uODIyLS42NzctMS41MDYtMS40ODktMS41MDZoLTE2LjgwOVYyOS43N2gxNi44MWMuODEgMCAxLjQ4OC0uNjE1IDEuNDg4LTEuNTA0di0zLjc2YzAtLjgyMy0uNjc3LTEuNTA2LTEuNDg5LTEuNTA2aC0xOC4yOTUtMy43MjVjLS44MTQgMC0xLjQ5LjY4My0xLjQ5IDEuNTA0djQxLjk5YzAgLjgyMi42NzYgMS41MDYgMS40OSAxLjUwNkgyMTYuNTFjLjgxMiAwIDEuNDg5LS42MTYgMS40ODktMS41MDV2LTMuNzYxek0zNDAgNjIuNzM0YzAtLjgyMS0uNjc3LTEuNTA1LTEuNDg5LTEuNTA1aC0xNi44MVY0OC44ODZoMTYuODFjLjgxIDAgMS40ODgtLjYxNSAxLjQ4OC0xLjUwNXYtMy43NmMwLS44MjItLjY3Ny0xLjUwNi0xLjQ4OS0xLjUwNmgtMTYuODA5VjI5Ljc3aDE2LjgxYy44MSAwIDEuNDg4LS42MTUgMS40ODgtMS41MDR2LTMuNzZjMC0uODIzLS42NzctMS41MDYtMS40ODktMS41MDZoLTE4LjI5NS0zLjcyNWMtLjgxNCAwLTEuNDkuNjgzLTEuNDkgMS41MDR2NDEuOTljMCAuODIyLjY3NiAxLjUwNiAxLjQ5IDEuNTA2SDMzOC41MWMuODEyIDAgMS40ODktLjYxNiAxLjQ4OS0xLjUwNXYtMy43NjF6TTQwNCA2Mi43MzRjMC0uODIxLS42NzctMS41MDUtMS40ODktMS41MDVoLTE2LjgxVjQ4Ljg4NmgxNi44MWMuODEyIDAgMS40ODgtLjYxNSAxLjQ4OC0xLjUwNXYtMy43NmMwLS44MjItLjY3Ny0xLjUwNi0xLjQ4OS0xLjUwNmgtMTYuODA5VjI5Ljc3aDE2LjgxYy44MSAwIDEuNDg4LS42MTUgMS40ODgtMS41MDR2LTMuNzZjMC0uODIzLS42NzctMS41MDYtMS40ODktMS41MDZoLTE4LjI5NS0zLjcyNWMtLjgxMyAwLTEuNDkuNjgzLTEuNDkgMS41MDR2NDEuOTljMCAuODIyLjY3NyAxLjUwNiAxLjQ5IDEuNTA2SDQwMi41MWMuODEyIDAgMS40ODktLjYxNiAxLjQ4OS0xLjUwNXYtMy43NjF6TTIyOS40NjcgMzQuNzcydjIuNTg2YzAgMS45NzQgMS44NjcgMy4xMyA0IDQuMDE1bDguODY2IDMuODFjMy4zMzMgMS40OTggNS42NjcgNC4xNSA1LjY2NyA4LjIzNHY0LjgzQzI0OCA2Mi4yNjQgMjM5LjczMiA2OSAyMzUuMzMyIDY5Yy0yLjk5OSAwLTcuNTMyLTIuMTc4LTExLjA2NC00LjE1MS0uODAxLS40MDgtMS4zMzQtMS4zNjItLjg2Ny0yLjQ1bDEuMTM0LTIuMzE0Yy40NjYtLjk1MiAxLjQ2NS0xLjE1NyAyLjMzMi0uNjggMi44NjYgMS40OTcgNi42NjYgMy4zMzQgOC40NjUgMy4zMzQgMiAwIDYuMjAxLTMuNDcgNi4yMDEtNS40NDR2LTIuOTkzYzAtMi4zMTQtMS43MzMtMy42MDctNC4xMzMtNC41NTlsLTguMjY3LTMuNTRjLTMuMi0xLjM2LTYuMTMzLTQuMTUtNi4xMzMtNy45NlYzMy44MmMwLTQuNDI0IDguNzMzLTEwLjgyIDEyLjkzNC0xMC44MiAyLjggMCA3LjMzMiAyLjExIDEwLjMzMyAzLjc0My45MzMuNDc3IDEuMiAxLjU2NS43OTggMi4zODFMMjQ2IDMxLjQzOGMtLjMzMy44MTYtMS40IDEuMDktMi4zMzIuNjgtMi4yNjctMS4wODgtNi4wNjctMi44NTctNy43MzMtMi44NTctMS45MzQgMC02LjQ2NyAzLjI2Ni02LjQ2NyA1LjUxMU0zNTkuNDUyIDI4Ljg4MmgtNi42NTd2MTUuMDQ2aDYuNzI1YzIuODE1IDAgNi4xMDktNC45OTIgNi4xMDktNy42NiAwLTIuNTk4LTMuNDMyLTcuMzg2LTYuMTc3LTcuMzg2bS0uMjc1IDIwLjkyN2gtNi4zODJ2MTYuNjg4YzAgLjg4OC0uNjE4IDEuNTAzLTEuNTggMS41MDNoLTMuNjM2Yy0uODI0IDAtMS41NzktLjYxNS0xLjU3OS0xLjUwM3YtNDAuMzVjMC0xLjU3MyAxLjg1NC0zLjE0NyAzLjM2My0zLjE0N2gxMC45ODJDMzY1LjM1NCAyMyAzNzIgMzAuNzg1IDM3MiAzNS45MTRjMCAzLjQyLTMuMDA4IDguNTYtNi4yMzMgMTAuNTQ0IDIuNjA3IDEuMzY3IDYuMTc2IDQuOTkyIDYuMTc2IDguMzQ0djExLjY5NWExLjUyIDEuNTIgMCAwIDEtMS41MSAxLjUwM2gtMy43NzRhMS41MiAxLjUyIDAgMCAxLTEuNTEtMS41MDN2LTEwLjY3YzAtMi4zOTQtMy4zNjItNi4wMTgtNS45NzItNi4wMThNMjY3LjA4OCAyOC44ODJoLTYuNDc3djE1LjA0Nmg2LjU0M2MyLjczOSAwIDUuNTYtNC45OTIgNS41Ni03LjY2IDAtMi41OTgtMi45NTYtNy4zODYtNS42MjYtNy4zODZtLS4yNjcgMjAuOTI3aC02LjIxdjE2LjY4OGMwIC44ODgtLjYwMiAxLjUwMy0xLjUzNiAxLjUwM2gtMy41NGMtLjggMC0xLjUzNS0uNjE1LTEuNTM1LTEuNTAzdi00MC4zNWMwLTEuNTczIDEuODAzLTMuMTQ3IDMuMjcyLTMuMTQ3aDEwLjY4NEMyNzIuODMgMjMgMjc5IDMxLjAyOCAyNzkgMzYuMTU2YzAgMy40Mi0yLjQ3MyA4LjAyLTUuMzMgMTAuNjgtMS40MDUgMS4zMDgtMy42MzYgMi45NzMtNi44NDkgMi45NzNNMTQyLjQ4NCA2OUMxMzguMTM2IDY5IDEzMCA2My4yMzcgMTMwIDU4LjI2M1YyNC44ODRjMC0uODI4LjcyMi0xLjUxNCAxLjUxNy0xLjUxNGgzLjQ4NWMuODUyIDAgMS41MS42ODYgMS41MSAxLjUxNFY1Ny4zN2MwIDIuNDE4IDMuOTMgNS4zNDkgNS45NzIgNS4zNDkgMi4wNDIgMCA1Ljk5OC0zLjMwMiA1Ljk5OC01LjcyVjI0LjUxNGMwLS44My41OTItMS41MTQgMS4zODgtMS41MTRoMy42MTNjLjc5NSAwIDEuNTE3LjY4NSAxLjUxNyAxLjUxNHYzMy4zOEMxNTUgNjIuNzE3IDE0Ni44MjcgNjkgMTQyLjQ4NCA2OU0xNzQuMTU3IDYyLjEyaC02LjM5NFY0Ny4wNzJoNi40NjJjMi44MDEgMCA2LjA4MSA0LjczIDYuMDgxIDcuMzk2IDAgMi42LTMuNDE2IDcuNjUtNi4xNDkgNy42NXptLTYuMzk0LTMzLjI0aDYuMzk0YzIuNzMzIDAgNi4xNDkgMy45MTggNi4xNDkgNi41MTggMCAyLjY2Ny0zLjI4IDcuMDIyLTYuMDgxIDcuMDIyaC02LjQ2MlYyOC44OHptMTIuOTg4IDE1Ljg2N2MzLjExMy0yLjA1MiA2LjI0OS02LjAzOSA2LjI0OS05LjM1QzE4NyAzMC4yNyAxODAuMDMyIDIzIDE3NS4wNDYgMjNoLTEwLjdjLTEuNTAxIDAtMy4zNDYgMS41NzMtMy4zNDYgMy4xNDV2NDAuMzUyYzAgLjg5Ljc1MSAxLjUwMyAxLjU3MSAxLjUwM2gxMi43MDZDMTgwLjI2MyA2OCAxODcgNTkuNiAxODcgNTQuNDdjMC0zLjMxLTMuMTM2LTcuNjQ3LTYuMjQ5LTkuNzIzek0xMDguOTA5IDQ1Ljc1bDE0LjcxNy0xNy4yNDRjLjUzLS42MjIuNTEyLTEuNTQtLjE1My0yLjEyMmwtMi44MTQtMi40NThhMS40ODggMS40ODggMCAwIDAtMi4wOTguMTU1bC0xMi44NyAxNS4wNzhWMjQuNTA1YzAtLjgyLS42MDgtMS41MDUtMS40ODYtMS41MDVoLTMuNzE4Yy0uODExIDAtMS40ODcuNjg0LTEuNDg3IDEuNTA1djQxLjk5MmMwIC44Mi42NzYgMS41MDMgMS40ODcgMS41MDNoMy43MThjLjg3OCAwIDEuNDg3LS42ODMgMS40ODctMS41MDNWNTIuMzM5bDEyLjg2OSAxNS4wOGExLjQ5IDEuNDkgMCAwIDAgMi4wOTguMTU0bDIuODE0LTIuNDU4Yy42NjUtLjU4Mi42ODMtMS41MDEuMTUzLTIuMTIxbC0xNC43MTctMTcuMjQ1ek00MjAuNTEgMjYuNDI1aC0xLjg0djIuNDM0aDEuODRjLjU5NCAwIDEuMjQ1LS40ODEgMS4yNDUtMS4xOSAwLS43NjMtLjY1LTEuMjQ0LTEuMjQ1LTEuMjQ0em0xLjEzMSA2LjAyOGwtMS43ODItMi43MTdoLTEuMTl2Mi43MTdoLS45NjF2LTYuODc3aDIuODAyYzEuMTYgMCAyLjIzNS44MiAyLjIzNSAyLjA5NCAwIDEuNTI5LTEuMzU4IDIuMDM4LTEuNzU0IDIuMDM4bDEuODQgMi43NDVoLTEuMTl6TTQyMCAyMy45MDZBNS4wNTYgNS4wNTYgMCAwIDAgNDE0LjkwNiAyOSA1LjA5IDUuMDkgMCAwIDAgNDIwIDM0LjA5NGE1LjA5IDUuMDkgMCAwIDAgNS4wOTQtNS4wOTNBNS4wNTYgNS4wNTYgMCAwIDAgNDIwIDIzLjkwNnpNNDIwIDM1Yy0zLjMxMSAwLTYtMi42ODktNi02IDAtMy4zMzkgMi42ODktNiA2LTYgMy4zNCAwIDYgMi42NjEgNiA2IDAgMy4zMTEtMi42NiA2LTYgNnoiLz48L2c+PHBhdGggZmlsbD0iIzAwQTk3MSIgZD0iTTY1IDcxLjY0N2wtMTktMTF2MjJ6TTY1IDE5LjY0N2wtMTktMTF2MjJ6TTE5LjY3OCA0NS42NDdMMzcgMzUuNTU2VjMuNjQ3TDEgMjQuNjJ2NDIuMDU1bDM2IDIwLjk3MlY1NS43Mzl6Ii8+PHBhdGggZmlsbD0iIzAwQTk3MSIgZD0iTTM3IDQ1LjY0N2wzNyAyMXYtNDJ6Ii8+PC9nPjwvc3ZnPg=="
															alt="KubeSphere®️"
														/>
													</p>
													<strong>告警消息</strong>
													<p class="line1">
														资源 {{.ResourceName}} {{.RuleName}} 告警解除
													</p>
													<p class="line2">首次发生时间：{{.FirstTime}}</p>
													<p>解除时间：{{.LastTime}}</p>
													<p>解除实时值：{{.LastValue}}</p>
													<hr />
													<p class="gray">
														* 此为系统邮件请勿回复
													</p>
												</td>
											</tr>
										</table>
									</td>
								</tr>
							</table>
	
							<div class="footer">
								<table
									role="presentation"
									border="0"
									cellpadding="0"
									cellspacing="0"
								>
									<tr>
										<td class="content-block">
											<span class="apple-link"
												>KubeSphere®️ 2019 All Rights Reserved.</span
											>
										</td>
									</tr>
								</table>
							</div>
						</div>
					</td>
					<td>&nbsp;</td>
				</tr>
			</table>
		</body>
	</html>
`
	EmailKubeSphereNotifyActiveTemplateEn = `
<!DOCTYPE html>
	<html>
		<head>
			<meta name="viewport" content="width=device-width" />
			<meta http-equiv="Content-Type" content="text/html; charset=UTF-8" />
			<title>Simple Transactional Email</title>
			<style>
				/* -------------------------------------
					GLOBAL RESETS
				------------------------------------- */
				/*All the styling goes here*/
				img {
					border: none;
					-ms-interpolation-mode: bicubic;
					max-width: 100%;
				}
	
				body {
					background-color: #ecf1f7;
					font-family: Roboto, PingFang SC, Lantinghei SC, Helvetica Neue,
						Helvetica, Arial, Microsoft YaHei, 微软雅黑, STHeitiSC-Light, simsun,
						宋体, WenQuanYi Zen Hei, WenQuanYi Micro Hei, sans-serif;
					-webkit-font-smoothing: antialiased;
					font-size: 14px;
					line-height: 1.4;
					margin: 0;
					padding: 0;
					-ms-text-size-adjust: 100%;
					-webkit-text-size-adjust: 100%;
					color: #242e42;
				}
	
				table {
					border-collapse: separate;
					mso-table-lspace: 0pt;
					mso-table-rspace: 0pt;
					width: 100%;
				}
				table td {
					font-size: 14px;
					vertical-align: top;
				}
	
				/* -------------------------------------
					BODY & CONTAINER
				------------------------------------- */
	
				.body {
					background-color: #eff0f5;
					width: 100%;
				}
	
				/* Set a max-width, and make it display as block so it will automatically stretch to that width, but will also shrink down on a phone or something */
				.container {
					display: block;
					margin: 0 auto !important;
					/* makes it centered */
					max-width: 780px;
					padding: 10px;
					padding-top: 80px;
					width: 780px;
				}
	
				/* This should also be a block element, so that it will fill 100% of the .container */
				.content {
					box-sizing: border-box;
					display: block;
					margin: 0 auto;
					max-width: 780px;
					padding: 10px;
				}
	
				/* -------------------------------------
					HEADER, FOOTER, MAIN
				------------------------------------- */
				.main {
					background: #ffffff;
					border-radius: 4px;
					border: solid 1px #e3e9ef;
					background-color: #ffffff;
				}
	
				.wrapper {
					box-sizing: border-box;
					padding: 48px;
				}
	
				.content-block {
					padding-bottom: 10px;
					padding-top: 10px;
				}
	
				.footer {
					clear: both;
					margin-top: 14px;
					text-align: center;
					width: 100%;
				}
				.footer td,
				.footer p,
				.footer span,
				.footer a {
					color: #8c96ad;
					font-size: 12px;
					text-align: center;
				}
				.gray {
					color: #8c96ad;
				}
				.logo {
					display: inline-block;
					vertical-align: middle;
				}
	
				/* -------------------------------------
					TYPOGRAPHY
				------------------------------------- */
				h1,
				h2,
				h3,
				h4 {
					color: #000000;
					font-weight: 400;
					line-height: 1.4;
					margin: 0;
					margin-bottom: 30px;
				}
	
				h1 {
					font-size: 35px;
					font-weight: 300;
					text-align: center;
					text-transform: capitalize;
				}
	
				p,
				ul,
				ol {
					font-size: 14px;
					font-weight: normal;
					line-height: 2;
					margin: 0;
					margin-bottom: 15px;
				}
				p li,
				ul li,
				ol li {
					list-style-position: inside;
					margin-left: 5px;
				}
	
				a {
					color: #329dce;
					font-weight: bold;
					text-decoration: none;
				}
	
				/* -------------------------------------
					BUTTONS
				------------------------------------- */
				.btn {
					box-sizing: border-box;
					width: 100%;
				}
				.btn > tbody > tr > td {
					padding-bottom: 15px;
				}
				.btn table {
					width: auto;
				}
				.btn table td {
					background-color: #ffffff;
					border-radius: 5px;
					text-align: center;
				}
				.btn a {
					background-color: #ffffff;
					border: solid 1px #3498db;
					border-radius: 5px;
					box-sizing: border-box;
					color: #3498db;
					cursor: pointer;
					display: inline-block;
					font-size: 14px;
					font-weight: bold;
					margin: 0;
					padding: 12px 25px;
					text-decoration: none;
					text-transform: capitalize;
				}
	
				.btn-primary table td {
					background-color: #3498db;
				}
	
				.btn-primary a {
					background-color: #3498db;
					border-color: #3498db;
					color: #ffffff;
				}
	
				/* -------------------------------------
					OTHER STYLES THAT MIGHT BE USEFUL
				------------------------------------- */
				.last {
					margin-bottom: 0;
				}
	
				.first {
					margin-top: 0;
				}
	
				.align-center {
					text-align: center;
				}
	
				.align-right {
					text-align: right;
				}
	
				.align-left {
					text-align: left;
				}
	
				.clear {
					clear: both;
				}
	
				.mt0 {
					margin-top: 0;
				}
	
				.mb0 {
					margin-bottom: 0;
				}
	
				.preheader {
					color: transparent;
					display: none;
					height: 0;
					max-height: 0;
					max-width: 0;
					opacity: 0;
					overflow: hidden;
					mso-hide: all;
					visibility: hidden;
					width: 0;
				}
	
				.powered-by a {
					text-decoration: none;
				}
	
				hr {
					border: 0;
					border-bottom: 1px solid #eff0f5;
					margin: 50px 0 12px;
				}
				.linkBtn {
					border-radius: 2px;
					box-shadow: 0 1px 3px 0 rgba(73, 33, 173, 0.16),
						0 1px 2px 0 rgba(52, 57, 69, 0.03);
					background-color: #242e42;
					color: #fff;
					line-height: 20px;
					padding: 8px 20px;
				}
				.link {
					font-size: 12px;
					font-style: normal;
					font-stretch: normal;
					line-height: 28px;
					letter-spacing: normal;
				}
				.platform {
					font-size: 14px;
					font-weight: 500;
					font-style: normal;
					font-stretch: normal;
					line-height: 20px;
					letter-spacing: normal;
					color: #343945;
					margin-left: 12px;
				}
				.line1 {
					margin-top: 22px;
					margin-bottom: 16px;
				}
				.line2 {
					line-height: 2;
					margin-top: 16px;
					margin-bottom: 20px;
				}
				.line3 {
					margin-bottom: 40px;
					margin-top: 16px;
				}
				.line4 {
					margin-bottom: 0px;
				}
				.line5 {
					margin-top: 0px;
				}
				.line6 {
					margin-bottom: 0px;
				}
	
				/* -------------------------------------
					RESPONSIVE AND MOBILE FRIENDLY STYLES
				------------------------------------- */
				@media only screen and (max-width: 620px) {
					table[class='body'] h1 {
						font-size: 28px !important;
						margin-bottom: 10px !important;
					}
					table[class='body'] p,
					table[class='body'] ul,
					table[class='body'] ol,
					table[class='body'] td,
					table[class='body'] span,
					table[class='body'] a {
						font-size: 16px !important;
					}
					table[class='body'] .wrapper,
					table[class='body'] .article {
						padding: 10px !important;
					}
					table[class='body'] .content {
						padding: 0 !important;
					}
					table[class='body'] .container {
						padding: 0 !important;
						width: 100% !important;
					}
					table[class='body'] .main {
						border-left-width: 0 !important;
						border-radius: 0 !important;
						border-right-width: 0 !important;
					}
					table[class='body'] .btn table {
						width: 100% !important;
					}
					table[class='body'] .btn a {
						width: 100% !important;
					}
					table[class='body'] .img-responsive {
						height: auto !important;
						max-width: 100% !important;
						width: auto !important;
					}
				}
	
				/* -------------------------------------
					PRESERVE THESE STYLES IN THE HEAD
				------------------------------------- */
				@media all {
					.ExternalClass {
						width: 100%;
					}
					.ExternalClass,
					.ExternalClass p,
					.ExternalClass span,
					.ExternalClass font,
					.ExternalClass td,
					.ExternalClass div {
						line-height: 100%;
					}
					.apple-link a {
						color: inherit !important;
						font-family: inherit !important;
						font-size: inherit !important;
						font-weight: inherit !important;
						line-height: inherit !important;
						text-decoration: none !important;
					}
					.btn-primary table td:hover {
						background-color: #34495e !important;
					}
					.btn-primary a:hover {
						background-color: #34495e !important;
						border-color: #34495e !important;
					}
				}
			</style>
		</head>
		<body class="">
			<span class="preheader"
				>This is preheader text. Some clients will show this text as a
				preview.</span
			>
			<table
				role="presentation"
				border="0"
				cellpadding="0"
				cellspacing="0"
				class="body"
			>
				<tr>
					<td>&nbsp;</td>
					<td class="container">
						<div class="content">
							<!-- START CENTERED WHITE CONTAINER -->
							<table role="presentation" class="main">
								<!-- START MAIN CONTENT AREA -->
								<tr>
									<td class="wrapper">
										<table
											role="presentation"
											border="0"
											cellpadding="0"
											cellspacing="0"
										>
											<tr>
												<td>
													<p styles="line-height: 30px;">
														<img
															width="150"
															height="30"
															src="data:image/svg+xml;base64,PHN2ZyB3aWR0aD0iNDI4IiBoZWlnaHQ9IjkwIiB4bWxucz0iaHR0cDovL3d3dy53My5vcmcvMjAwMC9zdmciPjxnIGZpbGw9Im5vbmUiIGZpbGwtcnVsZT0iZXZlbm9kZCI+PGcgZmlsbD0iIzQ0M0Q0RSI+PHBhdGggZD0iTTMwNy41MzQgMjNoLTMuNjY2Yy0uOCAwLTEuNDY4LjY4NC0xLjQ2OCAxLjUwNFY0Mi4wMUgyOTAuNlYyNC41MDRjMC0uODItLjYtMS41MDQtMS40NjYtMS41MDRoLTMuNjY4Yy0uOCAwLTEuNDY3LjY4NC0xLjQ2NyAxLjUwNHY0MS45OWMwIC44MjMuNjY3IDEuNTA2IDEuNDY3IDEuNTA2aDMuNjY4Yy44NjYgMCAxLjQ2Ni0uNjgzIDEuNDY2LTEuNTA1VjQ4Ljc4SDMwMi40djE3LjcxNWMwIC44MjIuNjY3IDEuNTA1IDEuNDY4IDEuNTA1aDMuNjY2Yy44NjcgMCAxLjQ2Ni0uNjgzIDEuNDY2LTEuNTA1di00MS45OWMwLS44Mi0uNi0xLjUwNS0xLjQ2Ni0xLjUwNU0yMTggNjIuNzM0YzAtLjgyMS0uNjc3LTEuNTA1LTEuNDg5LTEuNTA1aC0xNi44MVY0OC44ODZoMTYuODFjLjgxMiAwIDEuNDg4LS42MTUgMS40ODgtMS41MDV2LTMuNzZjMC0uODIyLS42NzctMS41MDYtMS40ODktMS41MDZoLTE2LjgwOVYyOS43N2gxNi44MWMuODEgMCAxLjQ4OC0uNjE1IDEuNDg4LTEuNTA0di0zLjc2YzAtLjgyMy0uNjc3LTEuNTA2LTEuNDg5LTEuNTA2aC0xOC4yOTUtMy43MjVjLS44MTQgMC0xLjQ5LjY4My0xLjQ5IDEuNTA0djQxLjk5YzAgLjgyMi42NzYgMS41MDYgMS40OSAxLjUwNkgyMTYuNTFjLjgxMiAwIDEuNDg5LS42MTYgMS40ODktMS41MDV2LTMuNzYxek0zNDAgNjIuNzM0YzAtLjgyMS0uNjc3LTEuNTA1LTEuNDg5LTEuNTA1aC0xNi44MVY0OC44ODZoMTYuODFjLjgxIDAgMS40ODgtLjYxNSAxLjQ4OC0xLjUwNXYtMy43NmMwLS44MjItLjY3Ny0xLjUwNi0xLjQ4OS0xLjUwNmgtMTYuODA5VjI5Ljc3aDE2LjgxYy44MSAwIDEuNDg4LS42MTUgMS40ODgtMS41MDR2LTMuNzZjMC0uODIzLS42NzctMS41MDYtMS40ODktMS41MDZoLTE4LjI5NS0zLjcyNWMtLjgxNCAwLTEuNDkuNjgzLTEuNDkgMS41MDR2NDEuOTljMCAuODIyLjY3NiAxLjUwNiAxLjQ5IDEuNTA2SDMzOC41MWMuODEyIDAgMS40ODktLjYxNiAxLjQ4OS0xLjUwNXYtMy43NjF6TTQwNCA2Mi43MzRjMC0uODIxLS42NzctMS41MDUtMS40ODktMS41MDVoLTE2LjgxVjQ4Ljg4NmgxNi44MWMuODEyIDAgMS40ODgtLjYxNSAxLjQ4OC0xLjUwNXYtMy43NmMwLS44MjItLjY3Ny0xLjUwNi0xLjQ4OS0xLjUwNmgtMTYuODA5VjI5Ljc3aDE2LjgxYy44MSAwIDEuNDg4LS42MTUgMS40ODgtMS41MDR2LTMuNzZjMC0uODIzLS42NzctMS41MDYtMS40ODktMS41MDZoLTE4LjI5NS0zLjcyNWMtLjgxMyAwLTEuNDkuNjgzLTEuNDkgMS41MDR2NDEuOTljMCAuODIyLjY3NyAxLjUwNiAxLjQ5IDEuNTA2SDQwMi41MWMuODEyIDAgMS40ODktLjYxNiAxLjQ4OS0xLjUwNXYtMy43NjF6TTIyOS40NjcgMzQuNzcydjIuNTg2YzAgMS45NzQgMS44NjcgMy4xMyA0IDQuMDE1bDguODY2IDMuODFjMy4zMzMgMS40OTggNS42NjcgNC4xNSA1LjY2NyA4LjIzNHY0LjgzQzI0OCA2Mi4yNjQgMjM5LjczMiA2OSAyMzUuMzMyIDY5Yy0yLjk5OSAwLTcuNTMyLTIuMTc4LTExLjA2NC00LjE1MS0uODAxLS40MDgtMS4zMzQtMS4zNjItLjg2Ny0yLjQ1bDEuMTM0LTIuMzE0Yy40NjYtLjk1MiAxLjQ2NS0xLjE1NyAyLjMzMi0uNjggMi44NjYgMS40OTcgNi42NjYgMy4zMzQgOC40NjUgMy4zMzQgMiAwIDYuMjAxLTMuNDcgNi4yMDEtNS40NDR2LTIuOTkzYzAtMi4zMTQtMS43MzMtMy42MDctNC4xMzMtNC41NTlsLTguMjY3LTMuNTRjLTMuMi0xLjM2LTYuMTMzLTQuMTUtNi4xMzMtNy45NlYzMy44MmMwLTQuNDI0IDguNzMzLTEwLjgyIDEyLjkzNC0xMC44MiAyLjggMCA3LjMzMiAyLjExIDEwLjMzMyAzLjc0My45MzMuNDc3IDEuMiAxLjU2NS43OTggMi4zODFMMjQ2IDMxLjQzOGMtLjMzMy44MTYtMS40IDEuMDktMi4zMzIuNjgtMi4yNjctMS4wODgtNi4wNjctMi44NTctNy43MzMtMi44NTctMS45MzQgMC02LjQ2NyAzLjI2Ni02LjQ2NyA1LjUxMU0zNTkuNDUyIDI4Ljg4MmgtNi42NTd2MTUuMDQ2aDYuNzI1YzIuODE1IDAgNi4xMDktNC45OTIgNi4xMDktNy42NiAwLTIuNTk4LTMuNDMyLTcuMzg2LTYuMTc3LTcuMzg2bS0uMjc1IDIwLjkyN2gtNi4zODJ2MTYuNjg4YzAgLjg4OC0uNjE4IDEuNTAzLTEuNTggMS41MDNoLTMuNjM2Yy0uODI0IDAtMS41NzktLjYxNS0xLjU3OS0xLjUwM3YtNDAuMzVjMC0xLjU3MyAxLjg1NC0zLjE0NyAzLjM2My0zLjE0N2gxMC45ODJDMzY1LjM1NCAyMyAzNzIgMzAuNzg1IDM3MiAzNS45MTRjMCAzLjQyLTMuMDA4IDguNTYtNi4yMzMgMTAuNTQ0IDIuNjA3IDEuMzY3IDYuMTc2IDQuOTkyIDYuMTc2IDguMzQ0djExLjY5NWExLjUyIDEuNTIgMCAwIDEtMS41MSAxLjUwM2gtMy43NzRhMS41MiAxLjUyIDAgMCAxLTEuNTEtMS41MDN2LTEwLjY3YzAtMi4zOTQtMy4zNjItNi4wMTgtNS45NzItNi4wMThNMjY3LjA4OCAyOC44ODJoLTYuNDc3djE1LjA0Nmg2LjU0M2MyLjczOSAwIDUuNTYtNC45OTIgNS41Ni03LjY2IDAtMi41OTgtMi45NTYtNy4zODYtNS42MjYtNy4zODZtLS4yNjcgMjAuOTI3aC02LjIxdjE2LjY4OGMwIC44ODgtLjYwMiAxLjUwMy0xLjUzNiAxLjUwM2gtMy41NGMtLjggMC0xLjUzNS0uNjE1LTEuNTM1LTEuNTAzdi00MC4zNWMwLTEuNTczIDEuODAzLTMuMTQ3IDMuMjcyLTMuMTQ3aDEwLjY4NEMyNzIuODMgMjMgMjc5IDMxLjAyOCAyNzkgMzYuMTU2YzAgMy40Mi0yLjQ3MyA4LjAyLTUuMzMgMTAuNjgtMS40MDUgMS4zMDgtMy42MzYgMi45NzMtNi44NDkgMi45NzNNMTQyLjQ4NCA2OUMxMzguMTM2IDY5IDEzMCA2My4yMzcgMTMwIDU4LjI2M1YyNC44ODRjMC0uODI4LjcyMi0xLjUxNCAxLjUxNy0xLjUxNGgzLjQ4NWMuODUyIDAgMS41MS42ODYgMS41MSAxLjUxNFY1Ny4zN2MwIDIuNDE4IDMuOTMgNS4zNDkgNS45NzIgNS4zNDkgMi4wNDIgMCA1Ljk5OC0zLjMwMiA1Ljk5OC01LjcyVjI0LjUxNGMwLS44My41OTItMS41MTQgMS4zODgtMS41MTRoMy42MTNjLjc5NSAwIDEuNTE3LjY4NSAxLjUxNyAxLjUxNHYzMy4zOEMxNTUgNjIuNzE3IDE0Ni44MjcgNjkgMTQyLjQ4NCA2OU0xNzQuMTU3IDYyLjEyaC02LjM5NFY0Ny4wNzJoNi40NjJjMi44MDEgMCA2LjA4MSA0LjczIDYuMDgxIDcuMzk2IDAgMi42LTMuNDE2IDcuNjUtNi4xNDkgNy42NXptLTYuMzk0LTMzLjI0aDYuMzk0YzIuNzMzIDAgNi4xNDkgMy45MTggNi4xNDkgNi41MTggMCAyLjY2Ny0zLjI4IDcuMDIyLTYuMDgxIDcuMDIyaC02LjQ2MlYyOC44OHptMTIuOTg4IDE1Ljg2N2MzLjExMy0yLjA1MiA2LjI0OS02LjAzOSA2LjI0OS05LjM1QzE4NyAzMC4yNyAxODAuMDMyIDIzIDE3NS4wNDYgMjNoLTEwLjdjLTEuNTAxIDAtMy4zNDYgMS41NzMtMy4zNDYgMy4xNDV2NDAuMzUyYzAgLjg5Ljc1MSAxLjUwMyAxLjU3MSAxLjUwM2gxMi43MDZDMTgwLjI2MyA2OCAxODcgNTkuNiAxODcgNTQuNDdjMC0zLjMxLTMuMTM2LTcuNjQ3LTYuMjQ5LTkuNzIzek0xMDguOTA5IDQ1Ljc1bDE0LjcxNy0xNy4yNDRjLjUzLS42MjIuNTEyLTEuNTQtLjE1My0yLjEyMmwtMi44MTQtMi40NThhMS40ODggMS40ODggMCAwIDAtMi4wOTguMTU1bC0xMi44NyAxNS4wNzhWMjQuNTA1YzAtLjgyLS42MDgtMS41MDUtMS40ODYtMS41MDVoLTMuNzE4Yy0uODExIDAtMS40ODcuNjg0LTEuNDg3IDEuNTA1djQxLjk5MmMwIC44Mi42NzYgMS41MDMgMS40ODcgMS41MDNoMy43MThjLjg3OCAwIDEuNDg3LS42ODMgMS40ODctMS41MDNWNTIuMzM5bDEyLjg2OSAxNS4wOGExLjQ5IDEuNDkgMCAwIDAgMi4wOTguMTU0bDIuODE0LTIuNDU4Yy42NjUtLjU4Mi42ODMtMS41MDEuMTUzLTIuMTIxbC0xNC43MTctMTcuMjQ1ek00MjAuNTEgMjYuNDI1aC0xLjg0djIuNDM0aDEuODRjLjU5NCAwIDEuMjQ1LS40ODEgMS4yNDUtMS4xOSAwLS43NjMtLjY1LTEuMjQ0LTEuMjQ1LTEuMjQ0em0xLjEzMSA2LjAyOGwtMS43ODItMi43MTdoLTEuMTl2Mi43MTdoLS45NjF2LTYuODc3aDIuODAyYzEuMTYgMCAyLjIzNS44MiAyLjIzNSAyLjA5NCAwIDEuNTI5LTEuMzU4IDIuMDM4LTEuNzU0IDIuMDM4bDEuODQgMi43NDVoLTEuMTl6TTQyMCAyMy45MDZBNS4wNTYgNS4wNTYgMCAwIDAgNDE0LjkwNiAyOSA1LjA5IDUuMDkgMCAwIDAgNDIwIDM0LjA5NGE1LjA5IDUuMDkgMCAwIDAgNS4wOTQtNS4wOTNBNS4wNTYgNS4wNTYgMCAwIDAgNDIwIDIzLjkwNnpNNDIwIDM1Yy0zLjMxMSAwLTYtMi42ODktNi02IDAtMy4zMzkgMi42ODktNiA2LTYgMy4zNCAwIDYgMi42NjEgNiA2IDAgMy4zMTEtMi42NiA2LTYgNnoiLz48L2c+PHBhdGggZmlsbD0iIzAwQTk3MSIgZD0iTTY1IDcxLjY0N2wtMTktMTF2MjJ6TTY1IDE5LjY0N2wtMTktMTF2MjJ6TTE5LjY3OCA0NS42NDdMMzcgMzUuNTU2VjMuNjQ3TDEgMjQuNjJ2NDIuMDU1bDM2IDIwLjk3MlY1NS43Mzl6Ii8+PHBhdGggZmlsbD0iIzAwQTk3MSIgZD0iTTM3IDQ1LjY0N2wzNyAyMXYtNDJ6Ii8+PC9nPjwvc3ZnPg=="
															alt="KubeSphere®️"
														/>
													</p>
													<strong>Alert Message</strong>
													<p class="line1">
														{{.ResourceName}} {{.RuleName}} has alerted
														{{.CumulatedCount}} time(s)
													</p>
													<p class="line2">First alert at: {{.FirstTime}}</p>
													<p>Last alert at: {{.LastTime}}</p>
													<p>The value on resume: {{.LastValue}}</p>
													<hr />
													<p class="gray">
														* Don't reply it, this is a system email
													</p>
												</td>
											</tr>
										</table>
									</td>
								</tr>
							</table>
	
							<div class="footer">
								<table
									role="presentation"
									border="0"
									cellpadding="0"
									cellspacing="0"
								>
									<tr>
										<td class="content-block">
											<span class="apple-link"
												>KubeSphere®️ 2019 All Rights Reserved.</span
											>
										</td>
									</tr>
								</table>
							</div>
						</div>
					</td>
					<td>&nbsp;</td>
				</tr>
			</table>
		</body>
	</html>
`
	EmailKubeSphereNotifyResumeTemplateEn = `
<!DOCTYPE html>
	<html>
		<head>
			<meta name="viewport" content="width=device-width" />
			<meta http-equiv="Content-Type" content="text/html; charset=UTF-8" />
			<title>Simple Transactional Email</title>
			<style>
				/* -------------------------------------
					GLOBAL RESETS
				------------------------------------- */
				/*All the styling goes here*/
				img {
					border: none;
					-ms-interpolation-mode: bicubic;
					max-width: 100%;
				}
	
				body {
					background-color: #ecf1f7;
					font-family: Roboto, PingFang SC, Lantinghei SC, Helvetica Neue,
						Helvetica, Arial, Microsoft YaHei, 微软雅黑, STHeitiSC-Light, simsun,
						宋体, WenQuanYi Zen Hei, WenQuanYi Micro Hei, sans-serif;
					-webkit-font-smoothing: antialiased;
					font-size: 14px;
					line-height: 1.4;
					margin: 0;
					padding: 0;
					-ms-text-size-adjust: 100%;
					-webkit-text-size-adjust: 100%;
					color: #242e42;
				}
	
				table {
					border-collapse: separate;
					mso-table-lspace: 0pt;
					mso-table-rspace: 0pt;
					width: 100%;
				}
				table td {
					font-size: 14px;
					vertical-align: top;
				}
	
				/* -------------------------------------
					BODY & CONTAINER
				------------------------------------- */
	
				.body {
					background-color: #eff0f5;
					width: 100%;
				}
	
				/* Set a max-width, and make it display as block so it will automatically stretch to that width, but will also shrink down on a phone or something */
				.container {
					display: block;
					margin: 0 auto !important;
					/* makes it centered */
					max-width: 780px;
					padding: 10px;
					padding-top: 80px;
					width: 780px;
				}
	
				/* This should also be a block element, so that it will fill 100% of the .container */
				.content {
					box-sizing: border-box;
					display: block;
					margin: 0 auto;
					max-width: 780px;
					padding: 10px;
				}
	
				/* -------------------------------------
					HEADER, FOOTER, MAIN
				------------------------------------- */
				.main {
					background: #ffffff;
					border-radius: 4px;
					border: solid 1px #e3e9ef;
					background-color: #ffffff;
				}
	
				.wrapper {
					box-sizing: border-box;
					padding: 48px;
				}
	
				.content-block {
					padding-bottom: 10px;
					padding-top: 10px;
				}
	
				.footer {
					clear: both;
					margin-top: 14px;
					text-align: center;
					width: 100%;
				}
				.footer td,
				.footer p,
				.footer span,
				.footer a {
					color: #8c96ad;
					font-size: 12px;
					text-align: center;
				}
				.gray {
					color: #8c96ad;
				}
				.logo {
					display: inline-block;
					vertical-align: middle;
				}
	
				/* -------------------------------------
					TYPOGRAPHY
				------------------------------------- */
				h1,
				h2,
				h3,
				h4 {
					color: #000000;
					font-weight: 400;
					line-height: 1.4;
					margin: 0;
					margin-bottom: 30px;
				}
	
				h1 {
					font-size: 35px;
					font-weight: 300;
					text-align: center;
					text-transform: capitalize;
				}
	
				p,
				ul,
				ol {
					font-size: 14px;
					font-weight: normal;
					line-height: 2;
					margin: 0;
					margin-bottom: 15px;
				}
				p li,
				ul li,
				ol li {
					list-style-position: inside;
					margin-left: 5px;
				}
	
				a {
					color: #329dce;
					font-weight: bold;
					text-decoration: none;
				}
	
				/* -------------------------------------
					BUTTONS
				------------------------------------- */
				.btn {
					box-sizing: border-box;
					width: 100%;
				}
				.btn > tbody > tr > td {
					padding-bottom: 15px;
				}
				.btn table {
					width: auto;
				}
				.btn table td {
					background-color: #ffffff;
					border-radius: 5px;
					text-align: center;
				}
				.btn a {
					background-color: #ffffff;
					border: solid 1px #3498db;
					border-radius: 5px;
					box-sizing: border-box;
					color: #3498db;
					cursor: pointer;
					display: inline-block;
					font-size: 14px;
					font-weight: bold;
					margin: 0;
					padding: 12px 25px;
					text-decoration: none;
					text-transform: capitalize;
				}
	
				.btn-primary table td {
					background-color: #3498db;
				}
	
				.btn-primary a {
					background-color: #3498db;
					border-color: #3498db;
					color: #ffffff;
				}
	
				/* -------------------------------------
					OTHER STYLES THAT MIGHT BE USEFUL
				------------------------------------- */
				.last {
					margin-bottom: 0;
				}
	
				.first {
					margin-top: 0;
				}
	
				.align-center {
					text-align: center;
				}
	
				.align-right {
					text-align: right;
				}
	
				.align-left {
					text-align: left;
				}
	
				.clear {
					clear: both;
				}
	
				.mt0 {
					margin-top: 0;
				}
	
				.mb0 {
					margin-bottom: 0;
				}
	
				.preheader {
					color: transparent;
					display: none;
					height: 0;
					max-height: 0;
					max-width: 0;
					opacity: 0;
					overflow: hidden;
					mso-hide: all;
					visibility: hidden;
					width: 0;
				}
	
				.powered-by a {
					text-decoration: none;
				}
	
				hr {
					border: 0;
					border-bottom: 1px solid #eff0f5;
					margin: 50px 0 12px;
				}
				.linkBtn {
					border-radius: 2px;
					box-shadow: 0 1px 3px 0 rgba(73, 33, 173, 0.16),
						0 1px 2px 0 rgba(52, 57, 69, 0.03);
					background-color: #242e42;
					color: #fff;
					line-height: 20px;
					padding: 8px 20px;
				}
				.link {
					font-size: 12px;
					font-style: normal;
					font-stretch: normal;
					line-height: 28px;
					letter-spacing: normal;
				}
				.platform {
					font-size: 14px;
					font-weight: 500;
					font-style: normal;
					font-stretch: normal;
					line-height: 20px;
					letter-spacing: normal;
					color: #343945;
					margin-left: 12px;
				}
				.line1 {
					margin-top: 22px;
					margin-bottom: 16px;
				}
				.line2 {
					line-height: 2;
					margin-top: 16px;
					margin-bottom: 20px;
				}
				.line3 {
					margin-bottom: 40px;
					margin-top: 16px;
				}
				.line4 {
					margin-bottom: 0px;
				}
				.line5 {
					margin-top: 0px;
				}
				.line6 {
					margin-bottom: 0px;
				}
	
				/* -------------------------------------
					RESPONSIVE AND MOBILE FRIENDLY STYLES
				------------------------------------- */
				@media only screen and (max-width: 620px) {
					table[class='body'] h1 {
						font-size: 28px !important;
						margin-bottom: 10px !important;
					}
					table[class='body'] p,
					table[class='body'] ul,
					table[class='body'] ol,
					table[class='body'] td,
					table[class='body'] span,
					table[class='body'] a {
						font-size: 16px !important;
					}
					table[class='body'] .wrapper,
					table[class='body'] .article {
						padding: 10px !important;
					}
					table[class='body'] .content {
						padding: 0 !important;
					}
					table[class='body'] .container {
						padding: 0 !important;
						width: 100% !important;
					}
					table[class='body'] .main {
						border-left-width: 0 !important;
						border-radius: 0 !important;
						border-right-width: 0 !important;
					}
					table[class='body'] .btn table {
						width: 100% !important;
					}
					table[class='body'] .btn a {
						width: 100% !important;
					}
					table[class='body'] .img-responsive {
						height: auto !important;
						max-width: 100% !important;
						width: auto !important;
					}
				}
	
				/* -------------------------------------
					PRESERVE THESE STYLES IN THE HEAD
				------------------------------------- */
				@media all {
					.ExternalClass {
						width: 100%;
					}
					.ExternalClass,
					.ExternalClass p,
					.ExternalClass span,
					.ExternalClass font,
					.ExternalClass td,
					.ExternalClass div {
						line-height: 100%;
					}
					.apple-link a {
						color: inherit !important;
						font-family: inherit !important;
						font-size: inherit !important;
						font-weight: inherit !important;
						line-height: inherit !important;
						text-decoration: none !important;
					}
					.btn-primary table td:hover {
						background-color: #34495e !important;
					}
					.btn-primary a:hover {
						background-color: #34495e !important;
						border-color: #34495e !important;
					}
				}
			</style>
		</head>
		<body class="">
			<span class="preheader"
				>This is preheader text. Some clients will show this text as a
				preview.</span
			>
			<table
				role="presentation"
				border="0"
				cellpadding="0"
				cellspacing="0"
				class="body"
			>
				<tr>
					<td>&nbsp;</td>
					<td class="container">
						<div class="content">
							<!-- START CENTERED WHITE CONTAINER -->
							<table role="presentation" class="main">
								<!-- START MAIN CONTENT AREA -->
								<tr>
									<td class="wrapper">
										<table
											role="presentation"
											border="0"
											cellpadding="0"
											cellspacing="0"
										>
											<tr>
												<td>
													<p styles="line-height: 30px;">
														<img
															width="150"
															height="30"
															src="data:image/svg+xml;base64,PHN2ZyB3aWR0aD0iNDI4IiBoZWlnaHQ9IjkwIiB4bWxucz0iaHR0cDovL3d3dy53My5vcmcvMjAwMC9zdmciPjxnIGZpbGw9Im5vbmUiIGZpbGwtcnVsZT0iZXZlbm9kZCI+PGcgZmlsbD0iIzQ0M0Q0RSI+PHBhdGggZD0iTTMwNy41MzQgMjNoLTMuNjY2Yy0uOCAwLTEuNDY4LjY4NC0xLjQ2OCAxLjUwNFY0Mi4wMUgyOTAuNlYyNC41MDRjMC0uODItLjYtMS41MDQtMS40NjYtMS41MDRoLTMuNjY4Yy0uOCAwLTEuNDY3LjY4NC0xLjQ2NyAxLjUwNHY0MS45OWMwIC44MjMuNjY3IDEuNTA2IDEuNDY3IDEuNTA2aDMuNjY4Yy44NjYgMCAxLjQ2Ni0uNjgzIDEuNDY2LTEuNTA1VjQ4Ljc4SDMwMi40djE3LjcxNWMwIC44MjIuNjY3IDEuNTA1IDEuNDY4IDEuNTA1aDMuNjY2Yy44NjcgMCAxLjQ2Ni0uNjgzIDEuNDY2LTEuNTA1di00MS45OWMwLS44Mi0uNi0xLjUwNS0xLjQ2Ni0xLjUwNU0yMTggNjIuNzM0YzAtLjgyMS0uNjc3LTEuNTA1LTEuNDg5LTEuNTA1aC0xNi44MVY0OC44ODZoMTYuODFjLjgxMiAwIDEuNDg4LS42MTUgMS40ODgtMS41MDV2LTMuNzZjMC0uODIyLS42NzctMS41MDYtMS40ODktMS41MDZoLTE2LjgwOVYyOS43N2gxNi44MWMuODEgMCAxLjQ4OC0uNjE1IDEuNDg4LTEuNTA0di0zLjc2YzAtLjgyMy0uNjc3LTEuNTA2LTEuNDg5LTEuNTA2aC0xOC4yOTUtMy43MjVjLS44MTQgMC0xLjQ5LjY4My0xLjQ5IDEuNTA0djQxLjk5YzAgLjgyMi42NzYgMS41MDYgMS40OSAxLjUwNkgyMTYuNTFjLjgxMiAwIDEuNDg5LS42MTYgMS40ODktMS41MDV2LTMuNzYxek0zNDAgNjIuNzM0YzAtLjgyMS0uNjc3LTEuNTA1LTEuNDg5LTEuNTA1aC0xNi44MVY0OC44ODZoMTYuODFjLjgxIDAgMS40ODgtLjYxNSAxLjQ4OC0xLjUwNXYtMy43NmMwLS44MjItLjY3Ny0xLjUwNi0xLjQ4OS0xLjUwNmgtMTYuODA5VjI5Ljc3aDE2LjgxYy44MSAwIDEuNDg4LS42MTUgMS40ODgtMS41MDR2LTMuNzZjMC0uODIzLS42NzctMS41MDYtMS40ODktMS41MDZoLTE4LjI5NS0zLjcyNWMtLjgxNCAwLTEuNDkuNjgzLTEuNDkgMS41MDR2NDEuOTljMCAuODIyLjY3NiAxLjUwNiAxLjQ5IDEuNTA2SDMzOC41MWMuODEyIDAgMS40ODktLjYxNiAxLjQ4OS0xLjUwNXYtMy43NjF6TTQwNCA2Mi43MzRjMC0uODIxLS42NzctMS41MDUtMS40ODktMS41MDVoLTE2LjgxVjQ4Ljg4NmgxNi44MWMuODEyIDAgMS40ODgtLjYxNSAxLjQ4OC0xLjUwNXYtMy43NmMwLS44MjItLjY3Ny0xLjUwNi0xLjQ4OS0xLjUwNmgtMTYuODA5VjI5Ljc3aDE2LjgxYy44MSAwIDEuNDg4LS42MTUgMS40ODgtMS41MDR2LTMuNzZjMC0uODIzLS42NzctMS41MDYtMS40ODktMS41MDZoLTE4LjI5NS0zLjcyNWMtLjgxMyAwLTEuNDkuNjgzLTEuNDkgMS41MDR2NDEuOTljMCAuODIyLjY3NyAxLjUwNiAxLjQ5IDEuNTA2SDQwMi41MWMuODEyIDAgMS40ODktLjYxNiAxLjQ4OS0xLjUwNXYtMy43NjF6TTIyOS40NjcgMzQuNzcydjIuNTg2YzAgMS45NzQgMS44NjcgMy4xMyA0IDQuMDE1bDguODY2IDMuODFjMy4zMzMgMS40OTggNS42NjcgNC4xNSA1LjY2NyA4LjIzNHY0LjgzQzI0OCA2Mi4yNjQgMjM5LjczMiA2OSAyMzUuMzMyIDY5Yy0yLjk5OSAwLTcuNTMyLTIuMTc4LTExLjA2NC00LjE1MS0uODAxLS40MDgtMS4zMzQtMS4zNjItLjg2Ny0yLjQ1bDEuMTM0LTIuMzE0Yy40NjYtLjk1MiAxLjQ2NS0xLjE1NyAyLjMzMi0uNjggMi44NjYgMS40OTcgNi42NjYgMy4zMzQgOC40NjUgMy4zMzQgMiAwIDYuMjAxLTMuNDcgNi4yMDEtNS40NDR2LTIuOTkzYzAtMi4zMTQtMS43MzMtMy42MDctNC4xMzMtNC41NTlsLTguMjY3LTMuNTRjLTMuMi0xLjM2LTYuMTMzLTQuMTUtNi4xMzMtNy45NlYzMy44MmMwLTQuNDI0IDguNzMzLTEwLjgyIDEyLjkzNC0xMC44MiAyLjggMCA3LjMzMiAyLjExIDEwLjMzMyAzLjc0My45MzMuNDc3IDEuMiAxLjU2NS43OTggMi4zODFMMjQ2IDMxLjQzOGMtLjMzMy44MTYtMS40IDEuMDktMi4zMzIuNjgtMi4yNjctMS4wODgtNi4wNjctMi44NTctNy43MzMtMi44NTctMS45MzQgMC02LjQ2NyAzLjI2Ni02LjQ2NyA1LjUxMU0zNTkuNDUyIDI4Ljg4MmgtNi42NTd2MTUuMDQ2aDYuNzI1YzIuODE1IDAgNi4xMDktNC45OTIgNi4xMDktNy42NiAwLTIuNTk4LTMuNDMyLTcuMzg2LTYuMTc3LTcuMzg2bS0uMjc1IDIwLjkyN2gtNi4zODJ2MTYuNjg4YzAgLjg4OC0uNjE4IDEuNTAzLTEuNTggMS41MDNoLTMuNjM2Yy0uODI0IDAtMS41NzktLjYxNS0xLjU3OS0xLjUwM3YtNDAuMzVjMC0xLjU3MyAxLjg1NC0zLjE0NyAzLjM2My0zLjE0N2gxMC45ODJDMzY1LjM1NCAyMyAzNzIgMzAuNzg1IDM3MiAzNS45MTRjMCAzLjQyLTMuMDA4IDguNTYtNi4yMzMgMTAuNTQ0IDIuNjA3IDEuMzY3IDYuMTc2IDQuOTkyIDYuMTc2IDguMzQ0djExLjY5NWExLjUyIDEuNTIgMCAwIDEtMS41MSAxLjUwM2gtMy43NzRhMS41MiAxLjUyIDAgMCAxLTEuNTEtMS41MDN2LTEwLjY3YzAtMi4zOTQtMy4zNjItNi4wMTgtNS45NzItNi4wMThNMjY3LjA4OCAyOC44ODJoLTYuNDc3djE1LjA0Nmg2LjU0M2MyLjczOSAwIDUuNTYtNC45OTIgNS41Ni03LjY2IDAtMi41OTgtMi45NTYtNy4zODYtNS42MjYtNy4zODZtLS4yNjcgMjAuOTI3aC02LjIxdjE2LjY4OGMwIC44ODgtLjYwMiAxLjUwMy0xLjUzNiAxLjUwM2gtMy41NGMtLjggMC0xLjUzNS0uNjE1LTEuNTM1LTEuNTAzdi00MC4zNWMwLTEuNTczIDEuODAzLTMuMTQ3IDMuMjcyLTMuMTQ3aDEwLjY4NEMyNzIuODMgMjMgMjc5IDMxLjAyOCAyNzkgMzYuMTU2YzAgMy40Mi0yLjQ3MyA4LjAyLTUuMzMgMTAuNjgtMS40MDUgMS4zMDgtMy42MzYgMi45NzMtNi44NDkgMi45NzNNMTQyLjQ4NCA2OUMxMzguMTM2IDY5IDEzMCA2My4yMzcgMTMwIDU4LjI2M1YyNC44ODRjMC0uODI4LjcyMi0xLjUxNCAxLjUxNy0xLjUxNGgzLjQ4NWMuODUyIDAgMS41MS42ODYgMS41MSAxLjUxNFY1Ny4zN2MwIDIuNDE4IDMuOTMgNS4zNDkgNS45NzIgNS4zNDkgMi4wNDIgMCA1Ljk5OC0zLjMwMiA1Ljk5OC01LjcyVjI0LjUxNGMwLS44My41OTItMS41MTQgMS4zODgtMS41MTRoMy42MTNjLjc5NSAwIDEuNTE3LjY4NSAxLjUxNyAxLjUxNHYzMy4zOEMxNTUgNjIuNzE3IDE0Ni44MjcgNjkgMTQyLjQ4NCA2OU0xNzQuMTU3IDYyLjEyaC02LjM5NFY0Ny4wNzJoNi40NjJjMi44MDEgMCA2LjA4MSA0LjczIDYuMDgxIDcuMzk2IDAgMi42LTMuNDE2IDcuNjUtNi4xNDkgNy42NXptLTYuMzk0LTMzLjI0aDYuMzk0YzIuNzMzIDAgNi4xNDkgMy45MTggNi4xNDkgNi41MTggMCAyLjY2Ny0zLjI4IDcuMDIyLTYuMDgxIDcuMDIyaC02LjQ2MlYyOC44OHptMTIuOTg4IDE1Ljg2N2MzLjExMy0yLjA1MiA2LjI0OS02LjAzOSA2LjI0OS05LjM1QzE4NyAzMC4yNyAxODAuMDMyIDIzIDE3NS4wNDYgMjNoLTEwLjdjLTEuNTAxIDAtMy4zNDYgMS41NzMtMy4zNDYgMy4xNDV2NDAuMzUyYzAgLjg5Ljc1MSAxLjUwMyAxLjU3MSAxLjUwM2gxMi43MDZDMTgwLjI2MyA2OCAxODcgNTkuNiAxODcgNTQuNDdjMC0zLjMxLTMuMTM2LTcuNjQ3LTYuMjQ5LTkuNzIzek0xMDguOTA5IDQ1Ljc1bDE0LjcxNy0xNy4yNDRjLjUzLS42MjIuNTEyLTEuNTQtLjE1My0yLjEyMmwtMi44MTQtMi40NThhMS40ODggMS40ODggMCAwIDAtMi4wOTguMTU1bC0xMi44NyAxNS4wNzhWMjQuNTA1YzAtLjgyLS42MDgtMS41MDUtMS40ODYtMS41MDVoLTMuNzE4Yy0uODExIDAtMS40ODcuNjg0LTEuNDg3IDEuNTA1djQxLjk5MmMwIC44Mi42NzYgMS41MDMgMS40ODcgMS41MDNoMy43MThjLjg3OCAwIDEuNDg3LS42ODMgMS40ODctMS41MDNWNTIuMzM5bDEyLjg2OSAxNS4wOGExLjQ5IDEuNDkgMCAwIDAgMi4wOTguMTU0bDIuODE0LTIuNDU4Yy42NjUtLjU4Mi42ODMtMS41MDEuMTUzLTIuMTIxbC0xNC43MTctMTcuMjQ1ek00MjAuNTEgMjYuNDI1aC0xLjg0djIuNDM0aDEuODRjLjU5NCAwIDEuMjQ1LS40ODEgMS4yNDUtMS4xOSAwLS43NjMtLjY1LTEuMjQ0LTEuMjQ1LTEuMjQ0em0xLjEzMSA2LjAyOGwtMS43ODItMi43MTdoLTEuMTl2Mi43MTdoLS45NjF2LTYuODc3aDIuODAyYzEuMTYgMCAyLjIzNS44MiAyLjIzNSAyLjA5NCAwIDEuNTI5LTEuMzU4IDIuMDM4LTEuNzU0IDIuMDM4bDEuODQgMi43NDVoLTEuMTl6TTQyMCAyMy45MDZBNS4wNTYgNS4wNTYgMCAwIDAgNDE0LjkwNiAyOSA1LjA5IDUuMDkgMCAwIDAgNDIwIDM0LjA5NGE1LjA5IDUuMDkgMCAwIDAgNS4wOTQtNS4wOTNBNS4wNTYgNS4wNTYgMCAwIDAgNDIwIDIzLjkwNnpNNDIwIDM1Yy0zLjMxMSAwLTYtMi42ODktNi02IDAtMy4zMzkgMi42ODktNiA2LTYgMy4zNCAwIDYgMi42NjEgNiA2IDAgMy4zMTEtMi42NiA2LTYgNnoiLz48L2c+PHBhdGggZmlsbD0iIzAwQTk3MSIgZD0iTTY1IDcxLjY0N2wtMTktMTF2MjJ6TTY1IDE5LjY0N2wtMTktMTF2MjJ6TTE5LjY3OCA0NS42NDdMMzcgMzUuNTU2VjMuNjQ3TDEgMjQuNjJ2NDIuMDU1bDM2IDIwLjk3MlY1NS43Mzl6Ii8+PHBhdGggZmlsbD0iIzAwQTk3MSIgZD0iTTM3IDQ1LjY0N2wzNyAyMXYtNDJ6Ii8+PC9nPjwvc3ZnPg=="
															alt="KubeSphere®️"
														/>
													</p>
													<strong>Alert Message</strong>
													<p class="line1">
														Alert {{.ResourceName}} {{.RuleName}} does not meet the threshold and has been resumed.
													</p>
													<p class="line2">First alert at: {{.FirstTime}}</p>
													<p>Last alert at: {{.LastTime}}</p>
													<p>The value on resume: {{.LastValue}}</p>
													<hr />
													<p class="gray">
														* Don't reply it, this is a system email
													</p>
												</td>
											</tr>
										</table>
									</td>
								</tr>
							</table>
	
							<div class="footer">
								<table
									role="presentation"
									border="0"
									cellpadding="0"
									cellspacing="0"
								>
									<tr>
										<td class="content-block">
											<span class="apple-link"
												>KubeSphere®️ 2019 All Rights Reserved.</span
											>
										</td>
									</tr>
								</table>
							</div>
						</div>
					</td>
					<td>&nbsp;</td>
				</tr>
			</table>
		</body>
	</html>
`
)
