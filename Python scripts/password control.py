print("voer uw wachtwoord in:")

def wachtwoord_controle():
    scorelengte = 0
    scorecijfers = 0
    scoreHoofdletters = 0
    scoreKleineletters = 0
    scoreSpecialetekens = 0

    wachtwoord = input()
    goedkeuring = True
    Specialetekens = ['!', '@', '#', '$', '%']

    if len(wachtwoord) < 8:
        print('lengte moet minimaal 8 zijn ')
        goedkeuring = False

    if len(wachtwoord) > 24:
        print('lengte moet niet groter zijn dan 20 ')
        goedkeuring = False

    if not any(char.isdigit() for char in wachtwoord):
        print('wachtwoord moet minimaal 1 cijfer hebben ')
        goedkeuring = False

    if not any(char.isupper() for char in wachtwoord):
        print('wachtwoord moet minimaal een hoofdletter hebben ')
        goedkeuring = False

    if not any(char.islower() for char in wachtwoord):
        print('wachtwoord moet minimaal een kleine letter hebben ')
        goedkeuring = False

    if not any(char in Specialetekens for char in wachtwoord):
        print('wachtwoord moet minimaal een speciaal teken hebben ')
        goedkeuring = False

    if goedkeuring == True:
        print("wachtwoord is geldig")
    if goedkeuring == False:
        print("wachtwoord is ongeldig!! ")

    # score van de lengte van het wachtwoord
    if len(wachtwoord) >= 9 and 11:
        scorelengte = 1
    if len(wachtwoord) >= 12 and 15:
        scorelengte = 2
    if len(wachtwoord) >= 16 and 19:
        scorelengte = 3
    if len(wachtwoord) == 20:
        scorelengte = 4
        if len(wachtwoord) >= 20:
            scorelengte = 0



    # score van de aantal numers in het wachtwoord
    if any(char.isdigit() for char in wachtwoord) >= 1 and 3:
        scorecijfers += 1
    if any(char.isdigit() for char in wachtwoord) >= 4 and 6:
        scorecijfers += 1
    if any(char.isdigit() for char in wachtwoord) >= 7 and 9:
        scorecijfers = 3

    # score van de aantal hoofdletters in het wachtwoord
    if any(char.isupper() for char in wachtwoord) >= 1 and 3:
        scoreHoofdletters = 1
    if any(char.isupper() for char in wachtwoord) >= 4 and 6:
        scoreHoofdletters = 2
    if any(char.isupper() for char in wachtwoord) >= 7 and 9:
        scoreHoofdletters = 3

    # score van de aantal kleineletters in het wachtwoord
    if any(char.islower() for char in wachtwoord) >= 1 and 3:
        scoreKleineletters = 1
    if any(char.islower() for char in wachtwoord) >= 4 and 6:
        scoreKleineletters = 2
    if any(char.islower() for char in wachtwoord) >= 7 and 9:
        scoreKleineletters = 3

    # score van de aantal specialetekens in het wachtwoord
    if any(char in Specialetekens for char in wachtwoord) >= 1 and 3:
        scoreSpecialetekens = 1
    if any(char in Specialetekens for char in wachtwoord) >= 4 and 6:
        scoreSpecialetekens = 2
    if any(char in Specialetekens for char in wachtwoord) >= 7 and 9:
        scoreSpecialetekens = 3


    print("score van lengte:", scorelengte)
    print("score van cijfers:", scorecijfers)
    print("score van hoofdletters:", scoreHoofdletters)
    print("score van kleineletters:", scoreKleineletters)
    print("score van specialetekens:", scoreSpecialetekens)
    print("het maximale aantal score is: 5")


wachtwoord_controle()