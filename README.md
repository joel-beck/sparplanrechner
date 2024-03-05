# Sparplanrechner 💰

Dieses Projekt enthält den Source Code für einen deutschsprachigen Sparplanrechner. Der Sparplanrechner kann kostenfrei über die Website [https://sparplanrechner.fly.dev](https://sparplanrechner.fly.dev) genutzt werden.

Nutzer können die Berechnung anhand der folgenden Parameter individualisieren:

-   Startkapital
-   Monatliche Sparrate
-   Erwartete jährliche Rendite
-   Anlagedauer
-   Erwartete jährliche Inflationsrate
-   Gewünschte jährliche Entnahmerate am Ende des Anlagezeitraums
-   Erwarteter Steuersatz auf die Kapitalerträge

Der Sparplanrechner ermittelt dann das erwartete Endkapital, jährliche und monatliche Entnahmebeträge (Brutto, Netto und Netto + inflationsbereinigt) sowie Zwischenbeträge für jedes Jahr des Anlagezeitraums.

## Motivation & Projektziele 🎯

Es existieren bereits zahlreiche deutschsprachige Sparplanrechner im Internet.
Zu nennen sind hier beispielsweise die hervorragenden Rechner von [Finanztip](https://www.finanztip.de/rechner/sparplanrechner/) sowie [Finanzfluss](https://www.finanzfluss.de/rechner/sparrechner/).

Keiner der bestehenden Lösungen vereint jedoch alle Features, die ich mir für einen Sparplanrechner wünsche.
Dazu zählen insbesondere die Berechnung inflationsbereinigter Endbeträge, Brutto- und Nettoentnahmebeträge basierend auf einer angestrebten jährlichen Entnahmerate sowie die Abgabe aller berechneten Zwischenbeträge für jedes Jahr des Anlagezeitraums.

Die [offensichtliche Lösung](https://xkcd.com/927/) bestand also darin, einen eigenen Sparplanrechner zu entwickeln.

Das Ziel dieses Projekts ist es, mittelfristig den besten kostenlosen deutschsprachigen Sparplanrechner zur Verfügung zu stellen.
Dazu sollen kontinuierlich neue Features hinzugefügt und die User Experience stetig verbessert werden.

## Development & Deployment 👨‍💻

Der Tech-Stack des Sparplanrechners besteht aus Go für das Backend mit dem [Templ](https://templ.guide) Paket für typsichere Templates, [Tailwind CSS](https://tailwindcss.com) und [Alpine JS](https://alpinejs.dev) für das Frontend und [HTMX](https://htmx.org/docs/) für die Kommunikation zwischen Frontend und Backend.

Die Web Applikation wird als Docker Container auf [Fly.io](https://fly.io) gehostet.
Bei jedem Merge auf den `main` Branch wird mit GitHub Actions automatisch ein neues Docker Image gebaut und auf Fly.io deployed. Zudem wird das Docker Image in die Container Registry Docker Hub gepusht.

## Contributions 🤝

Feedback, Bug Reports, Feature- und Pull Requests aus der Community sind jederzeit willkommen!
Dazu einfach ein Issue oder einen Pull Request gemeinsam mit einer kurzen Beschreibung auf dem GitHub Repository erstellen.

## English Description 🇺🇸🇬🇧

This project contains the source code for a German investment calculator.
The investment calculator can be accessed for free as a web application at [https://sparplanrechner.fly.dev](https://sparplanrechner.fly.dev).

Users can specify the starting capital, the monthly savings rate, the expected annual return, the investment period, the expected annual inflation rate, the desired annual takeout rate at the end of the investment period and the tax rate on the capital gains.

The investment calculator then calculates the expected final capital, annual and monthly takeout amounts (before tax, after tax as well as after tax + adjusted for inflation), and intermediate amounts for each year of the investment period.
