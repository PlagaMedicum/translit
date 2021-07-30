package translit

import "testing"

func TestCyrToLatin(t *testing.T) {
	tcs := []struct {
		name string
		in   string
		out  string
	}{
		{
			name: "Blok -- Nochj, ulica, fonarj, apteka",
			in: `Ночь, улица, фонарь, аптека,
Бессмысленный и тусклый свет.
Живи ещё хоть четверть века —
Всё будет так. Исхода нет.

Умрёшь — начнёшь опять сначала
И повторится всё, как встарь:
Ночь, ледяная рябь канала,
Аптека, улица, фонарь.

А. Блок`,
			out: `Nochj, ulica, fonarj, apteka,
Bessmyslennyj i tusklyj svet.
Zhivi jeshhio hotj chetvertj veka —
Vsio budet tak. Iskhoda net.

Umrioshj — nachnioshj opiatj snachala
I povtoritsia vsio, kak vstarj:
Nochj, ledianaja riabj kanala,
Apteka, ulica, fonarj.

A. Blok`,
		}, {
			name: "Rozhdestvenskij -- Abbrevjiatury",
			in: `Аббревиатуры

"Наша доля прекрасна, а воля — крепка!.."
РВС, ГОЭЛРО, ВЧК...

Наши марши взлетают до самых небес!
ЧТЗ, ГТО, МТС...

Кровь течёт на бетон из разорванных вен.
КПЗ, ЧСШ, ВМН...

Обожжённой, обугленной станет душа.
ПВО, РГК, ППШ...

Снова музыка в небе. Пора перемен.
АПК, ЭВМ, КВН...

"Наша доля прекрасна, а воля — крепка!"
SOS.
тчк

Р. Рождественский`,
			out: `Abbrevjiatury

"Nasha dolia prekrasna, a volia — krepka!.."
RVS, GOELRO, VChK...

Nashi marshi vzletajut do samykh nebes!
ChTZ, GTO, MTS...

Krovj techiot na beton iz razorvannykh ven.
KPZ, ChSSh, VMN...

Obozhzhionnoj, obuglennoj stanet dusha.
PVO, RGK, PPSh...

Snova muzyka v nebe. Pora peremen.
APK, EVM, KVN...

"Nasha dolia prekrasna, a volia — krepka!"
SOS.
tchk

R. Rozhdestvenskij`,
		},
	}

	for _, tc := range tcs {
		t.Run(tc.name, func(t *testing.T) {
			var res string
			res = CyrToLat(tc.in)
			if tc.out != res {
				t.Errorf("Convertion error.\nExpected: \"%s\"\nGot: \"%s\"", tc.out, res)
			}
		})
	}
}
