CREATE DATABASE 'class_note';

CREATE TABLE note
(
    id SERIAL PRIMARY KEY,
    created_time TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    username VARCHAR(10) NOT NULL,
    title VARCHAR(255) NOT NULL,
    content text
);

INSERT 
    INTO 
    note(username, title, content) 
VALUES
('SomeOne', 'Hello World', 'Hello World! Everybody'),
('SomeOne', 'Trash', 'waste material that is discarded by humans, usually due to a perceived lack of utility. The term generally does not encompass bodily waste products, purely liquid or gaseous wastes, nor toxic waste products. Garbage is commonly sorted and classified into kinds of material suitable for specific kinds of disposal.'),
('SomeOne', 'C#', 'C# is a general-purpose, multi-paradigm programming language. C# encompasses static typing, strong typing, lexically scoped, imperative, declarative, functional, generic, object-oriented (class-based), and component-oriented programming disciplines. C# was designed by Anders Hejlsberg from Microsoft in 2000 and was later approved as an international standard by Ecma (ECMA-334) in 2002 and ISO (ISO/IEC 23270) in 2003. Microsoft introduced C# along with .NET Framework and Visual Studio, both of which were closed-source. At the time, Microsoft had no open-source products. Four years later, in 2004, a free and open-source project called Mono began, providing a cross-platform compiler and runtime environment for the C# programming language. A decade later, Microsoft released Visual Studio Code (code editor), Roslyn (compiler), and the unified .NET platform (software framework), all of which support C# and are free, open-source, and cross-platform. Mono also joined Microsoft but was not merged into .NET. '),
('SomeOne', 'Go', 'Go is a statically typed, compiled programming language designed at Google by Robert Griesemer, Rob Pike, and Ken Thompson. Go is syntactically similar to C, but with memory safety, garbage collection, structural typing, and CSP-style concurrency. The language is often referred to as Golang because of its former domain name, golang.org, but the proper name is Go. There are two major implementations: Google self-hosting gc compiler toolchain targeting multiple operating systems, and WebAssembly. gofrontend, a frontend to other compilers, with the libgo library. With GCC the combination is gccgo; with LLVM the combination is gollvm.'),
('SomeOne', 'JavaScript', 'JavaScript often abbreviated JS, is a programming language that is one of the core technologies of the World Wide Web, alongside HTML and CSS. Over 97% of websites use Javascript on the client side for web page behavior, often incorporating third-party libraries. All major web browsers have a dedicated JavaScript engine to execute the code on the user device. JavaScript is a high-level, often just-in-time compiled, multi-paradigm language. It conforms to the ECMAScript specification, and has dynamic typing, prototype-based object-orientation, and first-class functions. It supports event-driven, functional, and imperative programming styles. It has application programming interfaces (APIs) for working with text, dates, regular expressions, standard data structures, and the Document Object Model (DOM).')
;

UPDATE
    note
SET
    title='C',
    content='C is a general-purpose, procedural computer programming language supporting structured programming, lexical variable scope, and recursion, with a static type system. By design, C provides constructs that map efficiently to typical machine instructions. It has found lasting use in applications previously coded in assembly language. Such applications include operating systems and various application software for computer architectures that range from supercomputers to PLCs and embedded systems.'
WHERE 
    title='Hello World'
;

DELETE
FROM   
    note
WHERE
    title='Trash'
;
