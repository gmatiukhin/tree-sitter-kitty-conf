import XCTest
import SwiftTreeSitter
import TreeSitterKittyconf

final class TreeSitterKittyconfTests: XCTestCase {
    func testCanLoadGrammar() throws {
        let parser = Parser()
        let language = Language(language: tree_sitter_kittyconf())
        XCTAssertNoThrow(try parser.setLanguage(language),
                         "Error loading Kittyconf grammar")
    }
}
