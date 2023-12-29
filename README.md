# Sparplanrechner 💰

Dieses Projekt enthält den Source Code für einen deutschsprachigen Sparplanrechner. Der Sparplanrechner kann kostenfrei über die Website [https://sparplanrechner.onrender.com](https://sparplanrechner.onrender.com) genutzt werden.

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

Das Frontend des Sparplanrechners wurde mit HTML, Tailwind CSS and Typescript entwickelt, das Backend ist in Go geschrieben.
Die Web Applikation wird als Docker Container auf [Render](https://render.com/) gehostet.
Bei jedem Pull Request auf den `main` Branch wird mit GitHub Actions automatisch ein neues Docker Image auf Docker Hub gepusht.
Das Update des Docker Images triggert anschließend ein neues Deployment auf Render.

## Contributions 🤝

Feedback, Bug Reports, Feature- und Pull Requests aus der Community sind jederzeit willkommen!
Dazu einfach ein Issue oder einen Pull Request gemeinsam mit einer kurzen Beschreibung auf dem GitHub Repository erstellen.

## English Description 🇺🇸🇬🇧

This project contains the source code for a German investment calculator.
The investment calculator can be accessed for free as a web application at [https://sparplanrechner.onrender.com](https://sparplanrechner.onrender.com).

Users can specify the starting capital, the monthly savings rate, the expected annual return, the investment period, the expected annual inflation rate, the desired annual takeout rate at the end of the investment period and the tax rate on the capital gains.

The investment calculator then calculates the expected final capital, annual and monthly takeout amounts (before tax, after tax as well as after tax + adjusted for inflation), and intermediate amounts for each year of the investment period.
